# 🚀 PocketURL - Enterprise URL Shortener

[![Go Version](https://img.shields.io/badge/Go-1.23-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![Supabase](https://img.shields.io/badge/Supabase-Database-3ECF8E?style=flat-square&logo=supabase)](https://supabase.com/)
[![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=flat-square&logo=docker)](https://www.docker.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)

**High-performance URL shortening service with automatic expiration, multi-domain support, and enterprise-grade architecture.**

## 🌟 Live Demo

🔗 **Visit:** [pocketurl.zip](https://pocketurl.zip)

## ✨ Key Features

- **⚡ Lightning Fast** - Sub-millisecond URL resolution
- **🕐 Auto-Expiration** - Automatic cleanup of expired links
- **🔒 HTTPS Security** - Production SSL/TLS encryption
- **🌐 Multi-Domain** - Custom domain support
- **🐳 Docker Ready** - Production containerization
- **📊 Smart Cleanup** - Background expired link removal

## 🛠️ Tech Stack

- **Backend:** Go 1.23 + Gin Framework
- **Database:** Supabase (PostgreSQL)
- **Infrastructure:** Docker + Alpine Linux
- **Libraries:** pgx/v5, scany, godotenv

## 🏗️ Architecture

```
pocketurl/
├── main.go              # Entry point
├── config/config.go     # Environment config
├── router/              # HTTP handlers & URL generation
├── persistence/         # Database operations & cleanup
└── Dockerfile           # Container build
```

## ⚡ Performance Features

- **Connection Pooling** - Optimized Supabase connections
- **6-Character Codes** - 56+ billion unique combinations
- **Background Cleanup** - Hourly expired link removal
- **CORS Support** - Cross-origin requests enabled

## 🚧 Future Enhancements

- [ ] Analytics dashboard with click tracking
- [ ] Custom vanity URLs
- [ ] Rate limiting and user authentication
- [ ] QR code generation
- [ ] Bulk URL import

## 📄 License

MIT License - see [LICENSE](LICENSE) for details.

---

⭐ **Star this repo if it helps you build better projects!**

Built with ❤️ in Go | Powered by Supabase