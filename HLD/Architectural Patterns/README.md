# Architectural Design Patterns 🏛️

Architectural design patterns are high-level strategies that concern the overall structure and organization of a software system. They provide a template for solving recurring design problems and help in creating scalable, maintainable, and efficient software architectures. 🔧🏗️

Types of Architectural Design Patterns:

- **Layered (n-tier) Architecture 🎂:** Organizes the system into layers with specific roles and responsibilities.
- **Client-Server Pattern 🖥️↔️📱:** Separates the system into two main components - the client (user interface) and the server (data storage and processing).
- **Master-Slave Pattern 👑👨‍👧‍👦:** Distributes work between a master component and several slave components.
- **Pipe-Filter Architecture 🚰🔍:** Processes data through a series of independent processing steps (filters) connected by pipes.
- **Broker Pattern 🤝:** Coordinates communication, such as requests, replies, and distributions, between objects.
- **Peer-to-Peer Pattern 🔄:** Distributes workloads among peers without the need for central coordination.
- **Event-Bus Pattern 🚌:** Handles communication between different parts of an application through events.
- **Model-View-Controller (MVC) 🧩:** Separates the application logic into three interconnected components.
- **Blackboard Pattern 📋:** Useful for problems for which no deterministic solution strategies are known.

These patterns help in creating robust, scalable, and maintainable software architectures by providing proven solutions to common architectural challenges. 🚀💡


# MVC Design Pattern 🧩

The Model-View-Controller (MVC) is an architectural pattern that separates an application into three main logical components: the Model, the View, and the Controller. Each of these components are built to handle specific development aspects of an application.

## Components:

- **Model:** Represents the data and business logic of the application.
- **View:** Displays the data to the user (UI).
- **Controller:** Handles the user input and mediates between Model and View.

## Advantages:

- Separation of Concerns: Clear division of responsibilities.
- Easier to Maintain: Changes in one component don't affect others.
- Parallel Development: Teams can work on different components simultaneously.
- Reusability: Components can be reused in other projects.

## Disadvantages:

- Complexity: Can be overkill for small applications.
- Learning Curve: Requires understanding of the pattern.
- Potential Performance Overhead: Due to the layered structure.

## Flow Chart:

![MVC Design Pattern](image-1.png)

In this flow:

1. User interacts with the Controller (e.g., clicks a button).
2. Controller updates the Model based on the user's action.
3. Model notifies the Controller of changes.
4. Controller selects appropriate View.
5. View displays updated information to the User.
6. Optionally, the Model may update the View directly in some implementations.


# Monolithic Architecture 🏢

A monolithic architecture is a traditional model where all components of an application are tightly integrated into a single, unified system. 🔒

## Advantages

- Simpler development process for small applications 👨‍💻
- Easier to deploy as a single unit 🚀
- Less complex inter-process communication 🔗

## Disadvantages

- Difficult to scale individual components 📈
- Changes in one part can affect the entire system 🔄
- Limited flexibility in technology choices 🔧

# Microservices Architecture 🧩

Microservices architecture is an approach where an application is built as a collection of small, independent services that communicate through APIs. 🌐

## Advantages

- Easier to scale individual services 🚀
- Greater flexibility in technology choices for each service 🛠️
- Improved fault isolation 🛡️
- Enables continuous deployment and faster updates 🔄

## Disadvantages

- Increased complexity in managing multiple services 🤯
- Potential for increased latency due to network communication 🐢
- More challenging to maintain data consistency across services 📊

Both architectures have their place in modern software development, with the choice depending on the specific needs and scale of the application. 🤔

# Converting a monolithic into microservices architecture

Converting a monolithic into microservices architecture requires careful evaluation of the business requirements. It involves different phases and each phase can have different patterns to create a micro services.

1. Decomposition

# Decomposition Pattern 🧩

Decomposition patterns are used to break down a monolithic application into smaller, more manageable services. There are three main approaches to decomposition:

1. **Domain-Driven Decomposition 🏢**: This pattern involves breaking down services based on business domains. It aligns the microservices structure with the organization's business capabilities.

    Example: In an e-commerce application 🛒, you might have separate microservices for user management, product catalog, order processing, and inventory management. Each of these represents a distinct business domain.

2. **Subdomain Decomposition 🧩**: This is a further refinement of domain-driven decomposition, where domains are divided into smaller subdomains. This allows for even more granular service definition.

    Example: Within the order processing domain 📦, you could further break it down into subdomains like order creation, payment processing, and shipping management.

3. **Functional Decomposition 🔧**: In this approach, services are separated based on business capabilities. Each microservice is responsible for a specific function or feature of the application.

    Example: For a social media platform 🤳, you might have microservices for user authentication, post creation and retrieval, friend management, and notification services. Each of these represents a specific function or feature of the application.

These decomposition patterns help in creating a more modular and scalable microservices architecture. They allow teams to develop, deploy, and scale services independently, improving overall system flexibility and maintainability. 🚀
