#!/bin/bash
set -e

CLUSTER_NAME="europe"
INVENTORY_DIR="../inventory/${CLUSTER_NAME}"
CONTROL_PLANE_INI="${INVENTORY_DIR}/control_panel.ini"
SSH_KEY="$HOME/.ssh/id_rsa"
SSH_USER="ubuntu"

echo "🛠️  Preparing environment..."
sudo apt install -y python3-pip

echo "📦 Installing dependencies..."
sudo apt update && sudo apt install -y python3-pip sshpass git
pip3 install --user ansible --break-system-packages # patch for the new debian version for root useage
export PATH="$HOME/.local/bin:$PATH"

echo "📥 Cloning Kubespray..."
git clone https://github.com/kubernetes-sigs/kubespray.git --depth 1 || true
cd kubespray
pip3 install -r requirements.txt --ignore-installed --break-system-packages # patch for the new debian version for root useage

echo "📂 Preparing control-plane inventory..."
cp -rfp inventory/sample inventory/${CLUSTER_NAME}
cp ../${CONTROL_PLANE_INI} inventory/${CLUSTER_NAME}/inventory.ini


pip3 install --user ansible-core==2.16.4 --break-system-packages # patch for the new debian version for root useage

echo "🚀 Deploying control-plane nodes..."
ansible-playbook -i inventory/${CLUSTER_NAME}/inventory.ini \
  --private-key=${SSH_KEY} \
  -u ${SSH_USER} \
  cluster.yml -b -v

echo "✅ Control plane deployed and kubeconfig downloaded!"
