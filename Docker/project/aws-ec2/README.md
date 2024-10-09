# Amazon Elastic Compute Cloud

> Amazon Elastic Compute Cloud (Amazon EC2) is a web service that provides resizable computing capacity—literally,
> servers in Amazon's data centers—that you use to build and host your software systems.
>
> [EC2 Documentation](https://docs.aws.amazon.com/ec2/)

For our project there are a few steps we need to take:

1. Create and launch EC2 instance, VPC (Virtual Private Cloud) and Security Group
2. Configure security group to expose all required ports to WWW
3. Connect to instance (SSH), install Docker and run our container

## Disadvantages of this Approach

1. We fully "own" the remote machine
    - Keep essential software updated
    - Manage network, security groups and firewall
2. SSH works but can be cumbersome if you deploy often

## Setting Up our EC2 Instance

Once you're logged in and registered in AWS, you should be on the AWS Management
Console. You should be able to search for available resources from there. Search
for the term EC2:

1. From the Dashboard, use the search bar to query for 'EC2'
   ![EC2 Search](../../../.attachments/EC2/EC2%20Search.png "EC2 Search")

2. Select it from the list and that should take you to the EC2 Dashboard
   ![EC2 Dashboard](../../../.attachments/EC2/EC2%20Dashboard.png "EC2 Dashbaord")

3. From the EC2 Dashboard go to "Launch Instance" panel and click on the "Launch Instance"
   button, that should bring you to the below screen
   ![Launch Instance](../../../.attachments/EC2/EC2%20LaunchInstance.png "EC2 Launch Instance")

4. For our use case we will set up our EC2 instance using the Amazon Linux AMI (Amazon Machine Image).
    - The description of our selected image should match:
      `Amazon Linux 2023 AMI 2023.4.20240319.1 x86_64 HVM kernel-6.1`
    - For the Instance Type, select any Free Tier Eligible option:
      `t2.micro` is the one I've chosen to select
    - Recommended: Create a new key-pair, to be used when SSHing into the instance creating it in the console
      will download a file. You get it once and are not able to download it again.

5. Leaving everything as default, ensure that a default VPC has been created and select and
   click on "Launch Instance"
   ![Launched Instance](../../.attachments/EC2/EC2%20LaunchedInstance.png "EC2 Launched Instance")

## Connecting to our EC2 Instance (using SSH)

1. On Linux or MacOS, the `ssh` command is available inside of your terminal
    - On `Windows > 10`, you may setup and
      use [Windows Subsystem for Linux](https://learn.microsoft.com/en-us/windows/wsl/install)
      or download an SSH Client like [PuTTY](https://putty.org/)

2. From the EC2 instances view (Step 5 of the above instructions) select your EC2 that you'd like to connect to.
3. Click "Connect" and this will bring you to a screen showing several options for connecting to a selected instance.
4. Click on the "SSH" tab and follow instructions for running the `ssh` command; it should be structured as such:
    ```bash
    ssh -i "EC2 .pem FILE PATH" <ec2-instance-dns>
    ```
5. You should then be connected via `ssh` to your EC2 instance with an intro screen:
    ```bash
       ,     #_
       ~\_  ####_        Amazon Linux 2023
      ~~  \_#####\
      ~~     \###|
      ~~       \#/ ___   https://aws.amazon.com/linux/amazon-linux-2023
       ~~       V~' '->
        ~~~         /
          ~~._.   _/
             _/ _/
           _/m/'
    [ec2-user@ip-xxx-xx-xx-xxx ~]$ 
    ```

## Installing Docker inside our EC2 Instance

1. To Install Docker onto our EC2 run the following commands:
    ```bash
    # ensure that system packages are updated
    sudo yum update -y
    sudo yum -y install docker

    sudo service docker start

    sudo usermod -a -G docker ec2-user
    ```

2. Log out and log back in, run this command to enable the Docker service
    ```bash
    sudo systemctl enable docker
    ```
3. Verify that the `docker` command may be run with:
    ```bash
    docker --version
    ```

