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


def bfs_seller(graph_neighbor: typing.Dict[str, typing.List[str]]) -> str:
    search_queue = deque()
    search_queue += graph_neighbor["you"]
    searched: typing.List[str] = []
    while search_queue:
        person = search_queue.popleft()
        pprint.pprint(person)
        if person not in searched:
            if is_seller(person):
                return person
            else:
                search_queue += graph_neighbor[person]
                searched.append(person)
    return str.__init__


def is_seller(name: str) -> bool:
    # NOTE(impcuong): The rightmost character from our friend's name must be letter "m"
    #   to indicate that they're a seller.
    return name[-1] == "m"


if __name__ == "__main__":
    seller = bfs_seller(init_graph())
    pprint.pprint("Seller: " + seller)
