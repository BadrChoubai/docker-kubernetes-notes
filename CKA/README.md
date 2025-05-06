# Certified Kubernetes Administrator Exam

These are my study notes I took in preparing for the Certified Kubernetes Administrator
Exam. To study, I was given a copy of the book by Chad M. Crowell: _Acing the Certified Kubernetes
Administrator Exam_, by a good friend of mine and wanted to make notes of what I read.

## How the Exam Is Structured

The Certified Kubernetes Administrator (CKA) certifican exam is a test designed to measure your competency
with Kubernetes, Those include:

1. Cluster architecture, installation, and configuration
2. Workdloads and scheduling
3. Services and networking
4. Storage
5. Troubleshooting

## What is a Kubernetes Administrator

As a Kubernetes Adminstrator, your primary role in an organization is to manage and maintain the reliability,
performance, and availability of a company's services. You'll may also be charged with identifying and monitoring
propers KPIs (key performance indicators) to ensure service health, minimizing MTTA (mean time to acknowledge), and
MTTR (mean time to repair).

> You are the one in-charge of understanding the inner workings of Kubernetes and how to translate that into value
> for end users

### What is Kubernetes?

Kubernetes is a piece of software that you interact with via a REST API. It's important to keep in mind
that Resources (a term used here to address objects inside of a cluster) can be managed through
the API exposed by it.

## Cluster Architecture

![Architecture Diagram](../.attachments/kubernetes-architecture-diagram.svg)

> [Read more about Kubernetes Architecture](../Kubernetes/ARCHITECTURE.md)
