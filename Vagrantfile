Vagrant.configure("2") do |config|

  config.vm.box = "ubuntu/xenial64"
  config.vm.network :private_network, ip: "10.10.10.10"

  config.vm.provision "ansible_local" do |ansible|
    ansible.verbose = false
    ansible.playbook = "provisioning/playbook.yml"
  end

end
