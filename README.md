# Alchemy at Lumen

This repository contains a collection of AI-powered tools and applications that demonstrate various use cases of OpenAI's language models and LangChain framework. The projects showcase different approaches to natural language processing, text transformation, and research capabilities.

## Projects Overview

| Project | Description | Key Features | Technologies |
|---------|-------------|--------------|--------------|
| [intro-chat-example](intro-chat-example/) | A text transformation demo that translates and modifies customer reviews | - Text translation and tone modification<br>- Multi-language support<br>- Professional tone transformation<br>- Streamlit web interface | - OpenAI API<br>- LangChain<br>- Streamlit<br>- Python |
| [wikipediatool](wikipediatool/) | A research tool that combines Wikipedia search with AI-powered analysis | - Wikipedia search integration<br>- Mathematical calculations<br>- Natural language processing<br>- CLI and Web interfaces | - OpenAI API<br>- LangChain<br>- Wikipedia API<br>- Streamlit |
| [weather-dashboard](https://github.com/anp0p/weather-dashboard) | A modern weather dashboard with AI-powered insights | - Real-time weather data fetching<br>- AI-powered weather commentary<br>- Durable workflow execution<br>- Clean, responsive web interface<br>- Fault-tolerant design | - Go<br>- Temporal<br>- OpenWeatherMap API<br>- OpenAI API |
| [chat](https://github.com/anp0p/chat) | A real-time chat application with multiple AI model support | - Multiple AI model integration (DeepSeek, LLaMA, Gemma)<br>- Real-time streaming responses<br>- Secure authentication<br>- Markdown support with syntax highlighting<br>- Dark/Light mode | - Next.js 14<br>- TypeScript<br>- Firebase<br>- Replicate API<br>- Tailwind CSS |
| [multi-model-image-generator](https://github.com/popand/multi-model-image-generator) | A modern AI image generation application | - Multiple AI models (Flux Pro, Flux Schnell, Ideogram)<br>- Automatic image saving and gallery<br>- Google authentication<br>- Real-time generation status<br>- Customizable parameters | - Next.js 14<br>- Firebase<br>- Replicate API<br>- Tailwind CSS<br>- TypeScript |
| [langgraph-tools-agent](https://github.com/anp0p/langgraph-tools-agent) | A Python-based tools agent with LangGraph workflow management | - Multiple tool support (Calculator, HTTP, Wikipedia, Code Execution)<br>- LangGraph-based workflow management<br>- Memory management for context<br>- Real-time streaming responses<br>- RESTful API interface | - Python<br>- LangGraph<br>- OpenAI GPT-4<br>- FastAPI |

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