
# Microservices Architecture 🧩

Microservices architecture is an approach where an application is built as a collection of small, independent services that communicate through APIs. 🌐

### Advantages

- Easier to scale individual services 🚀
- Greater flexibility in technology choices for each service 🛠️
- Improved fault isolation 🛡️
- Enables continuous deployment and faster updates 🔄

### Disadvantages

- Increased complexity in managing multiple services 🤯
- Potential for increased latency due to network communication 🐢
- More challenging to maintain data consistency across services 📊

Both architectures have their place in modern software development, with the choice depending on the specific needs and scale of the application. 🤔

# Converting a monolithic into microservices architecture

Converting a monolithic into microservices architecture requires careful evaluation of the business requirements. It involves different phases and each phase can have different patterns to create a micro services.

1. Decomposition
2. Refactoring
3. Database communication

# Decomposition Pattern 🧩

Decomposition patterns are used to break down a monolithic application into smaller, more manageable services. There are three main approaches to decomposition:

1. **Domain-Driven Decomposition 🏢**: This pattern involves breaking down services based on business domains. It aligns the microservices structure with the organization's business capabilities.

    Example: In an e-commerce application 🛒, you might have separate microservices for user management, product catalog, order processing, and inventory management. Each of these represents a distinct business domain.

2. **Subdomain Decomposition 🧩**: This is a further refinement of domain-driven decomposition, where domains are divided into smaller subdomains. This allows for even more granular service definition.

    Example: Within the order processing domain 📦, you could further break it down into subdomains like order creation, payment processing, and shipping management.

3. **Functional Decomposition 🔧**: In this approach, services are separated based on business capabilities. Each microservice is responsible for a specific function or feature of the application.

    Example: For a social media platform 🤳, you might have microservices for user authentication, post creation and retrieval, friend management, and notification services. Each of these represents a specific function or feature of the application.

These decomposition patterns help in creating a more modular and scalable microservices architecture. They allow teams to develop, deploy, and scale services independently, improving overall system flexibility and maintainability. 🚀

# Strangler Pattern

The Strangler pattern, also known as the Strangler Fig pattern 🌳, is a design pattern used in the gradual migration of legacy systems to microservices architecture. 

The Strangler pattern involves incrementally replacing specific pieces of functionality with new applications and services. The new system slowly grows around the old one, eventually replacing it entirely. This approach allows for a gradual transition, reducing risks associated with a complete rewrite. 🔄

Key aspects of the Strangler pattern:

- Gradual Migration: It allows for piece-by-piece migration rather than a complete overhaul. 🧩
- Risk Reduction: By replacing small parts at a time, it reduces the risk of system-wide failures. 🛡️
- Coexistence: The old and new systems can run side by side during the transition period. 🤝
- Incremental Value: Each replaced component can start providing value immediately. 📈

# Data Management in Microservices

In microservices architecture, data management can be approached in two main ways: database per service and shared database. Let's explore both approaches along with their advantages and disadvantages.

## 1. Database per Service 🏗️

In this approach, each microservice has its own dedicated database.

### Advantages
- Improved data encapsulation and isolation 🔒 
- Independent scalability of each service 🚀 
- Freedom to choose the most suitable database type for each service
- Easier to update and maintain individual services

### Disadvantages
- Complexity in managing distributed transactions or ACID properties across databases 🔗 - SAGA to rescue
- Challenging to perform queries across multiple services (Join queries) - CQRS to rescue 🔍 
- Increased storage requirements 💾 
- Higher operational costs 💰 

## 2. Shared Database 🤝

In this approach, multiple microservices share a common database.

### Advantages
- Easier data consistency management (ACID Properties)
- Simpler querying across services (Join Queries)
- Lower operational costs 💰 
- Faster development initially 🚀 

### Disadvantages
- Reduced data isolation 
- Potential for a monolithic database
- Schema changes can affect multiple services 🚧 
- Scalability limitations ⚖️ 

The choice between these approaches depends on factors such as the specific requirements of your microservices, scalability needs, data consistency requirements, and operational considerations. Some systems may even use a hybrid approach, combining both strategies where appropriate.

# SAGA Pattern 🧩

The SAGA pattern is a design pattern used in distributed systems to maintain data consistency across multiple services, each with its own database. It's particularly useful in microservices architectures where each service has its own database (DB per service approach). 🏗️

### How SAGA Maintains Consistency 🔄

SAGA breaks down a distributed transaction into a sequence of local transactions, each performed by a single service. If a step fails, the pattern executes compensating transactions to undo the changes made by the preceding steps. This ensures eventual consistency across all services. 🔁

## Types of SAGA 📊

## 1. Choreography-based SAGA 💃

In this approach, each service publishes domain events that trigger local transactions in other services.

### Advantages 🚀
- Simple to implement for simple workflows
- Loose coupling between services
- Good for simple scenarios with few steps

### Disadvantages 🤯
- Can be hard to understand and debug in complex scenarios 
- Risk of cyclic dependencies between services 🔄
- Difficult to implement for complex workflows 

Example: Online shopping system with three services: Order, Payment, and Inventory.
![Sequence of Events](image.png)
*Fig: Sequence of DB update events*

## 2. Orchestration-based SAGA 🎭

This approach uses a central orchestrator to manage the transaction and compensating actions.

### Advantages 🚀
- Easier to implement complex workflows 
- Centralized point of control and monitoring 
- Avoids cyclic dependencies between services 🚫🔄

### Disadvantages 🤯
- Introduces a single point of failure (the orchestrator) 💥
- Can lead to tighter coupling if not designed carefully 
- May become a performance bottleneck 

Example: Online shopping system with three services: Order, Payment, and Inventory.

![Events handled by Orchestrator](image-2.png)
*Fig: DB Events handled by Orchestrator*

In both cases, if any step fails (e.g., payment declined or out of stock), the SAGA pattern ensures that all previous steps are reversed, maintaining consistency across all services. 🔄✅