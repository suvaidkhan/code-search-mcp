# Code Search MCP

A high-performance semantic code search server built with Go that enables intelligent code discovery through natural language queries. Built on the Model Context Protocol (MCP), it combines AST parsing, vector embeddings, and real-time file monitoring to provide powerful code search capabilities.

## Features

- **🔍 Semantic Search**: Search your codebase using natural language queries powered by OpenAI embeddings
- **🌳 AST-Based Parsing**: Extracts functions, classes, and methods using Tree-sitter for accurate code structure understanding
- **⚡ Real-Time Monitoring**: Automatic re-indexing on file changes using fsnotify with concurrent processing
- **🎯 Targeted Retrieval**: Reduces token usage by retrieving only relevant code segments
- **🔧 Git Integration**: Respects `.gitignore` patterns to exclude unnecessary files
- **🗄️ Vector Storage**: Fast similarity search using chromem-go for efficient vector database operations
- **🌐 Multi-Language Support**: Supports Go, Python, TypeScript, JavaScript, and Markdown

## How It Works

1. **File Discovery**: Scans your project directory while respecting `.gitignore` rules
2. **AST Parsing**: Uses Tree-sitter to parse source files and extract code entities (functions, classes, methods)
3. **Embedding Generation**: Creates semantic embeddings using OpenAI's API
4. **Vector Storage**: Stores embeddings in chromem-go for fast similarity search
5. **Real-Time Updates**: Monitors file changes and automatically re-indexes modified files
6. **Semantic Search**: Queries return the most relevant code segments based on semantic similarity

## Prerequisites

- Go 1.21 or higher
- OpenAI API key
- Git (for gitignore integration)

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/code-search-mcp.git
cd code-search-mcp

# Install dependencies
go mod download

# Build the server
go build -o code-search-mcp
```

## Configuration

Set your OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY=your-api-key-here
```

## Usage

### Starting the Server

```bash
./code-search-mcp --path /path/to/your/project
```

### Command Line Options

```
--path         Path to the project directory to index (required)
--port         Server port (default: 8080)
--watch        Enable file watching for real-time updates (default: true)
--languages    Comma-separated list of languages to index (default: all supported)
```

### Example Queries

Once the server is running, you can search your codebase:

```bash
# Find authentication logic
curl -X POST http://localhost:8080/search \
  -d '{"query": "user authentication and login"}'

# Find database connection code
curl -X POST http://localhost:8080/search \
  -d '{"query": "database connection setup"}'

# Find error handling patterns
curl -X POST http://localhost:8080/search \
  -d '{"query": "error handling middleware"}'
```

## Supported Languages

- **Go** (.go)
- **Python** (.py)
- **TypeScript** (.ts, .tsx)
- **JavaScript** (.js, .jsx)
- **Markdown** (.md)

Additional language support can be added by extending the Tree-sitter grammar integration.

## Architecture

```
┌─────────────────┐
│  File Watcher   │
│   (fsnotify)    │
└────────┬────────┘
         │
         ▼
┌─────────────────┐     ┌──────────────────┐
│  AST Parser     │────▶│  Code Extractor  │
│  (Tree-sitter)  │     │                  │
└─────────────────┘     └────────┬─────────┘
                                 │
                                 ▼
                        ┌──────────────────┐
                        │ OpenAI Embeddings│
                        │       API        │
                        └────────┬─────────┘
                                 │
                                 ▼
                        ┌──────────────────┐
                        │  Vector Database │
                        │   (chromem-go)   │
                        └────────┬─────────┘
                                 │
                                 ▼
                        ┌──────────────────┐
                        │  Search Engine   │
                        │   (Similarity)   │
                        └──────────────────┘
```

## Performance

- **Concurrent Processing**: File monitoring and indexing run in parallel
- **Incremental Updates**: Only changed files are re-indexed
- **Efficient Storage**: Vector database optimized for similarity search
- **Token Optimization**: Returns only relevant code segments, reducing context size

## Development

### Running Tests

```bash
go test ./...
```

### Building from Source

```bash
go build -o code-search-mcp ./cmd/server
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Tree-sitter](https://tree-sitter.github.io/) for powerful AST parsing
- [chromem-go](https://github.com/philippgille/chromem-go) for efficient vector storage
- [OpenAI](https://openai.com/) for embedding generation
- [fsnotify](https://github.com/fsnotify/fsnotify) for file system monitoring
- [MCP](https://modelcontextprotocol.io/) for the protocol specification

## Roadmap

- [ ] Add support for more programming languages
- [ ] Implement caching layer for frequently accessed embeddings
- [ ] Add web UI for interactive search
- [ ] Support for local embedding models (Ollama, etc.)
- [ ] Multi-repository indexing
- [ ] Advanced filtering (by file type, date, author)
- [ ] Export search results to various formats

## Support

If you encounter any issues or have questions, please [open an issue](https://github.com/yourusername/code-search-mcp/issues) on GitHu
