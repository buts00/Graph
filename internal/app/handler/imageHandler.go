package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func (h *Handler) sendImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Error retrieving file from request: "+err.Error())
		return
	}

	// Open the file
	uploadedFile, err := file.Open()
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Error opening file: "+err.Error())
		return
	}
	defer uploadedFile.Close()

	// Create a new multipart writer
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create a new form file field and copy the file content into it
	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Error creating form file: "+err.Error())
		return
	}
	_, err = io.Copy(part, uploadedFile)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Error copying file content: "+err.Error())
		return
	}

	// Close the multipart writer to finalize the request body
	err = writer.Close()
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Error closing multipart writer: "+err.Error())
		return
	}

	// Send the multipart request to the Flask application
	resp, err := http.Post("http://localhost:14880", writer.FormDataContentType(), &requestBody)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Error sending request to Flask application: "+err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Unexpected status code from Flask application")
		return
	}

	// Read the response body from Flask
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Error reading response from Flask application: "+err.Error())
		return
	}

	fmt.Println(string(responseBytes))
	ctx.String(http.StatusOK, string(responseBytes))
}
