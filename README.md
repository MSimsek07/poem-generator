# Poem Generator

This project is a web application that generates poems based on user-provided prompts using the OpenAI GPT-3.5-turbo model. The backend is implemented in Go, and the frontend is built with Streamlit.

## Features

- Generate poems based on user prompts.
- Uses OpenAI's GPT-3.5-turbo model for poem generation.
- Simple and intuitive Streamlit interface.

## Prerequisites

- Go 1.22.4 or later
- Python 3.7 or later
- OpenAI API key

## Setup

### Backend (Go)

1. **Clone the repository**:

    ```sh
    git clone https://github.com/MSimsek07/poem-generator.git
    cd poem-generator
    ```

2. **Initialize Go module**:

    ```sh
    go mod init poem-generator
    go mod tidy
    ```

3. **Set your OpenAI API key**:

    ```sh
    export OPENAI_API_KEY=your_openai_api_key
    ```

4. **Build and run the Go server**:

    ```sh
    go build -o poem-generator .
    ./poem-generator
    ```

    The server will run on `http://localhost:8000`.

### Frontend (Streamlit)

1. **Install Streamlit**:

    ```sh
    pip install streamlit
    ```

2. **Create `streamlit_app.py`** with the given content


3. **Run the Streamlit app**:

    ```sh
    streamlit run streamlit_app.py
    ```

## Usage

1. Open your web browser and go to `http://localhost:8501`.
2. Enter a prompt in the text input box.
3. Click the "Generate Poem" button.
4. The generated poem will be displayed on the page.

## Project Structure

- `main.go`: The Go backend server code.
- `streamlit_app.py`: The Streamlit frontend code.
- `go.mod`: The Go module file.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any changes you would like to make.

## License

This project is licensed under the MIT License.
