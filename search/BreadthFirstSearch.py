import pprint
import typing
from collections import deque


def init_graph() -> typing.Dict[str, typing.List[str]]:
    graph = {}
    graph["you"] = ["alice", "bob", "claire"]
    graph["bob"] = ["anuj", "peggy"]
    graph["alice"] = ["peggy"]
    graph["claire"] = ["thom", "jonny"]
    graph["anuj"] = []
    graph["peggy"] = []
    graph["thom"] = []
    graph["jonny"] = []
    return graph


def add_vertex(
    graph: typing.Dict[str, typing.List[str]], key: str, value: typing.List[str]
) -> None:
    for k in graph.keys():
        if k == key:
            return
    graph[key] = [value if isinstance(value, list) else None]


def bfs_seller(graph_neighbor: typing.Dict[str, typing.List[str]]) -> str:
    search_queue = deque()
    search_queue += graph_neighbor["you"]
    searched: typing.List[str] = []
    while search_queue:
        person: str = search_queue.popleft()
        if is_seller(person):
            return person
        if person not in searched and not is_seller(person):
            pprint.pprint(person)
            search_queue += graph_neighbor[person]
            searched.append(person)
    return str.__init__


def is_seller(name: str) -> bool:
    # NOTE(impcuong): The rightmost character from our friend's name must be letter "m"
    #   to indicate that they're a seller.
    return name[-1] == "m"


def find_lowest_cost_node(
    costs_graph: typing.Dict[str, int], processed: typing.List[str]
) -> str:
    lowest_cost = float("inf")
    lowest_cost_node = None
    for node in costs_graph:
        cost = costs_graph[node]
        if cost < lowest_cost and node not in processed:
            lowest_cost = cost
            lowest_cost_node = node
    return lowest_cost_node


if __name__ == "__main__":
    seller = bfs_seller(init_graph())
    pprint.pprint("Seller: " + seller)

    routine_graph = {}
    add_vertex(routine_graph, "Wake-up", ["Shower", "Brush-teeth"])
    add_vertex(routine_graph, "Brush-teeth", ["Breakfast"])
    pprint.pprint(routine_graph)

    # NOTE(impcuong): Dijkstra’s algorithm.
    weighted_graph = {}
    weighted_graph["start"] = {}
    weighted_graph["start"]["a"] = 6
    weighted_graph["start"]["b"] = 2

    weighted_graph["a"] = {}
    weighted_graph["a"]["fin"] = 1

    weighted_graph["b"] = {}
    weighted_graph["b"]["a"] = 3
    weighted_graph["b"]["fin"] = 5

    weighted_graph["fin"] = {}
    pprint.pprint(weighted_graph)

    collectMapKeyFn = weighted_graph["start"].keys
    pprint.pprint(collectMapKeyFn())

    costs_graph = {}
    costs_graph["a"] = 6
    costs_graph["b"] = 2
    costs_graph["fin"] = float("inf")
    pprint.pprint(costs_graph)

    parent_graph = {}
    parent_graph["a"] = "start"
    parent_graph["b"] = "start"
    parent_graph["fin"] = None
    pprint.pprint(parent_graph)

    processed_nodes = []
    node = find_lowest_cost_node(costs_graph, processed=processed_nodes)
    pprint.pprint(node)
    while node is not None:
        cost = costs_graph[node]
        neighbors = weighted_graph[node]
        for k in neighbors.keys():
            new_cost = cost + neighbors[k]
            if costs_graph[k] > new_cost:
                costs_graph[k] = new_cost
                parent_graph[k] = node
        processed_nodes.append(node)
        node = find_lowest_cost_node(costs_graph, processed=processed_nodes)
        pprint.pprint(node)

    pprint.pprint(parent_graph)
    pprint.pprint(costs_graph)
