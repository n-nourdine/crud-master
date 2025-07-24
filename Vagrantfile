Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/focal64"

  # Gateway VM
  config.vm.define "gateway-vm" do |gateway|
    gateway.vm.network "forwarded_port", guest: 8080, host: 8080
    gateway.vm.provision "shell", path: "scripts/install_common.sh"
    gateway.vm.provision "shell", path: "scripts/install_gateway.sh"
    gateway.vm.provision "file", source: ".env", destination: "/home/vagrant/.env"
    gateway.vm.provision "file", source: "src/api-gateway", destination: "/home/vagrant/api-gateway"
  end

  # Inventory VM
  config.vm.define "inventory-vm" do |inventory|
    inventory.vm.network "forwarded_port", guest: 8081, host: 8081
    inventory.vm.network "forwarded_port", guest: 5432, host: 5432
    inventory.vm.provision "shell", path: "scripts/install_common.sh"
    inventory.vm.provision "shell", path: "scripts/setup_postgres.sh"
    inventory.vm.provision "shell", path: "scripts/install_inventory.sh"
    inventory.vm.provision "file", source: ".env", destination: "/home/vagrant/.env"
    inventory.vm.provision "file", source: "src/inventory-app", destination: "/home/vagrant/inventory-app"
  end

  # Billing VM
  config.vm.define "billing-vm" do |billing|
    billing.vm.network "forwarded_port", guest: 8082, host: 8082
    billing.vm.network "forwarded_port", guest: 5672, host: 5672
    billing.vm.network "forwarded_port", guest: 15672, host: 15672 # RabbitMQ UI
    billing.vm.provision "shell", path: "scripts/install_common.sh"
    billing.vm.provision "shell", path: "scripts/setup_postgres.sh"
    billing.vm.provision "shell", path: "scripts/setup_rabbitmq.sh"
    billing.vm.provision "shell", path: "scripts/install_billing.sh"
    billing.vm.provision "file", source: ".env", destination: "/home/vagrant/.env"
    billing.vm.provision "file", source: "src/billing-app", destination: "/home/vagrant/billing-app"
  end
end