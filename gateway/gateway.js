const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require('@apollo/gateway');

// Initialize an ApolloGateway instance and pass it an array of
// your subgraph names and URLs
const gateway = new ApolloGateway({
  serviceList: [
    { name: 'user-go', url: 'http://localhost:8080/query' },
    { name: 'post-go', url: 'http://localhost:8081/query' },
  ],
});

// Pass the ApolloGateway to the ApolloServer constructor
const server = new ApolloServer({
  gateway,

  // Disable subscriptions (not currently supported with ApolloGateway)
  subscriptions: false,
});

server.listen().then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});