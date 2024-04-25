# Alpine Linux tabanlı Python image
FROM python:alpine

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy requirements.txt file to the container's workspace
COPY requirements.txt ./

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the local package files to the container's workspace
COPY . .

# Command to run the application
CMD ["python", "app.py"]
