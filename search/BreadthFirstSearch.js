let weightedGraph = {};
weightedGraph["start"] = {};
weightedGraph["start"]["a"] = 5;
weightedGraph["start"]["b"] = 2;

weightedGraph["a"] = {};
weightedGraph["a"]["c"] = 4;
weightedGraph["a"]["d"] = 2;

weightedGraph["b"] = {};
weightedGraph["b"]["a"] = 8;
weightedGraph["b"]["d"] = 7;

weightedGraph["c"] = {};
weightedGraph["c"]["d"] = 6;
weightedGraph["c"]["fin"] = 3;

weightedGraph["d"] = {};
weightedGraph["d"]["fin"] = 1;

weightedGraph["fin"] = {};

let costGraph = {};
costGraph["a"] = 5;
costGraph["b"] = 2;
costGraph["c"] = Number.MAX_VALUE;
costGraph["d"] = Number.MAX_VALUE;
costGraph["fin"] = Number.MAX_VALUE;

let parentGraph = {};
parentGraph["a"] = "start";
parentGraph["b"] = "start";
parentGraph["c"] = "";
parentGraph["d"] = "";
parentGraph["fin"] = "";

let processedNodes = [];
function findLowestCostNode(costs, processed) {
  if (!(processed instanceof Array)) return "".toString();
  let lowestCost = Number.MAX_VALUE;
  let lowestCostNode = "";
  Object.keys(costs)
    .forEach((nodeName) => {
      if (!processed.includes(nodeName) && costs[nodeName] < lowestCost) {
        lowestCost = costs[nodeName];
        lowestCostNode = nodeName;
      }
    })
  return lowestCostNode;
}

let curNode = findLowestCostNode(costGraph, processedNodes);
while (curNode !== "") {
  const cost = costGraph[curNode];
  const neighbors = weightedGraph[curNode];
  Object.keys(neighbors).forEach((neighborName) => {
    const newCost = cost + neighbors[neighborName] /* neighborWeight */;
    if (costGraph[neighborName] > newCost) {
      costGraph[neighborName] = newCost;
      parentGraph[neighborName] = curNode;
    }
  });
  processedNodes.push(curNode);
  curNode = findLowestCostNode(costGraph, processedNodes);
}

console.log("Processed: ", processedNodes);
console.log("Parents:   ", parentGraph);
console.log("Costs:     ", costGraph);
