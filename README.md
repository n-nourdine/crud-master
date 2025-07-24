# CRUD Master

A microservices-based movie streaming platform with an API Gateway, Inventory API (PostgreSQL), and Billing API (RabbitMQ). Built entirely in Go, running on three Vagrant-managed VMs.

## Stack
- **Language**: Go 1.21
- **API Gateway**: Routes HTTP requests and RabbitMQ messages
- **Inventory API**: RESTful CRUD API with PostgreSQL (`movies_db`)
- **Billing API**: Consumes RabbitMQ messages and stores in PostgreSQL (`billing_db`)
- **Database**: PostgreSQL
- **Message Queue**: RabbitMQ
- **Virtualization**: Vagrant (Ubuntu 20.04 VMs)
- **Process Manager**: PM2
- **Testing**: Postman
- **Documentation**: OpenAPI (Swagger)

## Setup Instructions

1. **Install Dependencies**:
   - Install [Vagrant](https://www.vagrantup.com/downloads) and [VirtualBox](https://www.virtualbox.org/wiki/Downloads).
   - Install [Go](https://golang.org/dl/) for local development (optional, as VMs install Go).
   - Install [Postman](https://www.postman.com/downloads/) for testing.

2. **Clone Repository**:
   ```bash
   git clone <repository-url>
   cd crud-master
3. **Configure Environment**:

    Copy `.env.example` to `.env` and update credentials if needed (defaults provided for testing).

4. **Start VMs**:
    ```bash
    vagrant up
    ```
    This creates three VMs: `gateway-vm` (port 8080), `inventory-vm` (port 8081, PostgreSQL 5432), and `billing-vm` (port 8082, RabbitMQ 5672/15672).


5. **Access VMs (optional)**:
    ```bash
    vagrant ssh gateway-vm
    vagrant ssh inventory-vm
    vagrant ssh billing-vm
6. **Manage Services with PM2: Inside a VM**:
    ````bash
    sudo pm2 list
    sudo pm2 stop <app_name>
    sudo pm2 start <app_name>