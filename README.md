# Alchemy at Lumen

This repository contains a collection of AI-powered tools and applications that demonstrate various use cases of OpenAI's language models and LangChain framework. The projects showcase different approaches to natural language processing, text transformation, and research capabilities.

## Projects Overview

| Project | Description | Key Features | Technologies |
|---------|-------------|--------------|--------------|
| [intro-chat-example](intro-chat-example/) | A text transformation demo that translates and modifies customer reviews | - Text translation and tone modification<br>- Multi-language support<br>- Professional tone transformation<br>- Streamlit web interface | - OpenAI API<br>- LangChain<br>- Streamlit<br>- Python |
| [wikipediatool](wikipediatool/) | A research tool that combines Wikipedia search with AI-powered analysis | - Wikipedia search integration<br>- Mathematical calculations<br>- Natural language processing<br>- CLI and Web interfaces | - OpenAI API<br>- LangChain<br>- Wikipedia API<br>- Streamlit |

## Getting Started

### Prerequisites

- Python 3.x
- OpenAI API key
- Required Python packages (listed in each project's requirements.txt)

### OpenAI API Key Setup

To use these applications, you'll need an OpenAI API key:

1. Create an account at [OpenAI](https://openai.com)
2. Navigate to the API section in your dashboard
3. Create a new API key
4. Set up billing information
5. Securely store your API key

Each project has its own method for API key configuration:
- `intro-chat-example`: Uses a `.env` file
- `wikipediatool`: Uses `apikey.py`

## Project Details

### Intro Chat Example

A demonstration of text transformation capabilities using OpenAI's language models. It includes:
- Direct OpenAI API implementation
- LangChain integration
- Streamlit web interface for customer review transformation
- Multi-language support and tone modification

[View detailed documentation](intro-chat-example/README.md)

### Wikipedia Tool

A research assistant that combines Wikipedia's knowledge base with AI-powered analysis. Features:
- Command-line interface for research queries
- Web interface for user-friendly interaction
- Integration with Wikipedia search
- Mathematical calculation capabilities

[View detailed documentation](wikipediatool/README.md)

## Important Note

These projects are demonstrative in nature and should not be considered production-ready. They serve as examples of how to integrate and utilize OpenAI's language models and LangChain framework in various applications.

## License

This project is licensed under the MIT License - see the LICENSE file in each project directory for details. 