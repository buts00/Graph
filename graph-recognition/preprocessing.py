import cv2 as cv
import numpy as np

WIDTH_LIMIT = 1280
HEIGHT_LIMIT = 800
MIN_BRIGHTNESS = 110

def preprocess(image):
    reshaped = reshape(image, WIDTH_LIMIT, HEIGHT_LIMIT)
    
    gray = cv.cvtColor(reshaped, cv.COLOR_BGR2GRAY)
    
    binary = threshold(gray, MIN_BRIGHTNESS)

    reshaped, binary = crop_background(reshaped,binary)
    
    return reshaped, binary

def reshape(image, width_lim, height_lim):
    height, width = image.shape[:2]
    if height > width:
        image = cv.rotate(image, cv.ROTATE_90_CLOCKWISE)
    
    scale = min(width_lim / width, height_lim / height)
    image = cv.resize(image, None, fx=scale, fy=scale)
    return image

def threshold(gray_image, min_brightness):
    thresh_type = cv.THRESH_BINARY if np.mean(gray_image) < min_brightness else cv.THRESH_BINARY_INV
    _, binary = cv.threshold(gray_image, 0, 255, thresh_type + cv.THRESH_OTSU)
    return binary

def crop_background(reshaped,binary):
    
    contours, _ = cv.findContours(binary, cv.RETR_EXTERNAL, cv.CHAIN_APPROX_SIMPLE)
    x_min, y_min = np.inf, np.inf
    x_max, y_max = -np.inf, -np.inf

    for contour in contours:
        x, y, w, h = cv.boundingRect(contour)
        x_min = min(x_min, x)
        y_min = min(y_min, y)
        x_max = max(x_max, x + w)
        y_max = max(y_max, y + h)

    if contours:
        cropped_images = [img[y_min-10:y_max+10, x_min-10:x_max+10] for img in [reshaped] + [binary]]
        return cropped_images[0], cropped_images[1]
    else:
        return reshaped,binary

