# Capstone Project: AWS Elastic Kubernetes Service

For the capstone project in the course, we focused on moving away from our local development environment (`minikube`) and start to learn
about the workflow, deployment options, etc. involved in deploying software to a production cloud environment: AWS (Amazon Web
Services).

> ## EKS vs ECS
>
>  | **AWS EKS (Elastic Kubernetes Service)**             | **AWS ECS (Elastic Container Service)**     |
>  |:-----------------------------------------------------|:--------------------------------------------|
>  | Managed service for Kubernetes deployments           | Managed service for Container deployments   |
>  | No AWS-specific syntax or philosophy required        | AWS-specific syntax and philosophy applies  |
>  | Use standard Kubernetes configurations and resources | Use AWS-specific configuration and concepts |

**Sections**:

- [Stretch Goals](#stretch-goals)
  - [OpenTofu](#opentofu)

---

## Stretch Goals

For myself, I wanted to add two stretch goals:

1. Build the applications in a different programming language (Go)
2. Use an infrastructure-as-code tool to manage resources deployed to AWS (OpenTofu)

### OpenTofu

OpenTofu is an open-source fork of Terraform managed by the Linux Foundation. From
the [Manifesto](https://opentofu.org/manifesto/),
it was created in response to Hashicorp's decision to change the license on the Terraform source code from MPL (Mozilla
Public License) to a non-open source license.

The maintainers of OpenTofu and its users believe that the license change ultimately harms the open-source community and
the ecosystem that Terraform developed over the nine years leading up to the change.

> **Install OpenTofu on Ubuntu using `snap`**
>
> ```shell
> snap install --classic opentofu 
> ```


