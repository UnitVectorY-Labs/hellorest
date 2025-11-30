[![GitHub release](https://img.shields.io/github/release/UnitVectorY-Labs/hellorest.svg)](https://github.com/UnitVectorY-Labs/hellorest/releases/latest) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![Active](https://img.shields.io/badge/Status-Active-green)](https://guide.unitvectorylabs.com/bestpractices/status/#active)

# hellorest

A trivial Go API returning `{"hello": "world"}` on a `GET /` request. This project is designed as a simple Docker container to test API responses and configurations.

All other verbs and paths return a 404.

## Purpose

This Docker image provides a minimal REST API for testing. It can be used to validate deployment setups, networking, or any environment that needs a basic HTTP server with predictable behavior.

This is intended to be deployed as a place holder to test various API and microservice components.
