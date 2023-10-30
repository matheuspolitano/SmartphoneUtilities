CREATE TABLE SmartphoneUtilities (
    utility_id INT PRIMARY KEY,
    utility_name VARCHAR(255) NOT NULL,
    brand VARCHAR(50) NOT NULL,
    model_compatibility VARCHAR(255),
    price DECIMAL(10, 2) NOT NULL,
    quantity_in_stock INT DEFAULT 0,
    description TEXT,
    image_url VARCHAR(500),
    date_added DATE NOT NULL DEFAULT CURRENT_DATE
);
