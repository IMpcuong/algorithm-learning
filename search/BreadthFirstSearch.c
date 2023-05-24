#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#define MAX_CAPACITY 4096

// Resource: https://www.digitalocean.com/community/tutorials/hash-table-in-c-plus-plus.

/*
  Command:
  ```
  gcc -Wall -Wextra -pedantic \
    -Wshadow -Wformat=2 -Wconversion \
    -Wpointer-arith -Wredundant-decls -Wnested-externs \
    -Wno-long-long -Wcast-qual -ggdb \
    -std=c11 -O2 \
    -o out BreadthFirstSearch.c
  ```
*/

uint32_t hasher_operation(char *str)
{
  uint32_t idx = 0;
  for (int j = 0; j < MAX_CAPACITY; j++)
    idx += (unsigned char)(str[j] << __CHAR_BIT__);
  idx = (idx << 13) ^ idx;
  return ((idx * (idx * idx * 15731 + 789221) + 1376312589) & 0x7fffffff);
}

typedef struct HashItem
{
  char *key;
  char *value;
} HashItem;

HashItem *create_hash_item(char *key, char *value)
{
  HashItem *item_ptr;
  item_ptr = (HashItem *)malloc(sizeof *item_ptr);
  item_ptr->key = (char *)(strlen(key) + 1);
  item_ptr->value = (char *)(strlen(value) + 1);
  strcpy(item_ptr->key, key);
  strcpy(item_ptr->value, value);
  return item_ptr;
}

void free_hash_item(HashItem *item_ptr)
{
  free(item_ptr->key);
  free(item_ptr->value);
  free(item_ptr);
}

typedef struct HashMap
{
  HashItem **items;
  uint16_t size;
  uint16_t count;
} HashMap;

HashMap *create_hash_map(uint16_t size)
{
  HashMap *map_ptr;
  map_ptr = (HashMap *)malloc(sizeof *map_ptr);
  map_ptr->size = size;
  map_ptr->count = 0;

  map_ptr->items = (HashItem **)calloc(map_ptr->size, sizeof(HashItem *));
  for (uint16_t i = 0; i < map_ptr->size; i++)
    map_ptr->items[i] = NULL;
  return map_ptr;
}

void insert_into_hash_map(HashMap *map_ptr, char *key, char *value)
{
  HashItem *item = create_hash_item(key, value);
  uint32_t item_idx = hasher_operation(key);
  HashItem *cur_map_item = map_ptr->items[item_idx];
  if (cur_map_item == NULL)
  {
    if (map_ptr->size == map_ptr->count)
    {
      // NOTE: All slots have been allocated.
      printf("ERROR: Hash Map is full\n");
      free_hash_item(item);
      return;
    }
    map_ptr->items[item_idx] = item;
    map_ptr->count++;
  }
  else
  {
    // NOTE: `{ key: value }` has already existed, to address this,
    //    must update the item value to the new one.
    if (strcmp(cur_map_item->key, key) == 0)
    {
      strcpy(map_ptr->items[item_idx]->value, value);
      return;
    }
    else
    {
      // FIXME(impcuong): Implements later.
      // handle_collision(map_ptr, item);
      return;
    }
  }
}

HashItem *search_item_hash_map(HashMap *map_ptr, char *key)
{
  uint32_t item_idx = hasher_operation(key);
  HashItem *item = map_ptr->items[item_idx];
  if (item == NULL)
    if (strcmp(item->key, key) == 0)
      return item;
  return NULL;
}

void print_search_operation(HashMap *map_ptr, char *key)
{
  HashItem *item;
  if ((item = search_item_hash_map(map_ptr, key)) == NULL)
  {
    printf("ERROR: [Key: %s] does not exist\n", key);
    return;
  }
  else
  {
    printf("[Key: %s], [Value: %s]\n", key, item->value);
  }
}

void delete_item_from_hash_map(HashMap *map_ptr, char *key)
{
  uint32_t item_idx = hasher_operation(key);
  HashItem *item = map_ptr->items[item_idx];
  if (item == NULL)
    return;

  if (strcmp(item->key, key) == 0)
  {
    map_ptr->items[item_idx] = NULL;
    free_hash_item(item);
    map_ptr->count--;
    return;
  }
}

void free_hash_map(HashMap *map_ptr)
{
  for (uint16_t i = 0; i < map_ptr->size; i++)
  {
    HashItem *item_ptr = map_ptr->items[i];
    if (item_ptr != NULL)
      free_hash_item(item_ptr);
  }
  free(map_ptr->items);
  free(map_ptr);
}

void print_hash_map(HashMap *map_ptr)
{
  printf("\n*** Hash Map ***\n-------------------\n");
  for (int i = 0; i < map_ptr->size; i++)
    if (map_ptr->items[i])
      printf("Index:%d, Key:%s, Value:%s\n", i, map_ptr->items[i]->key, map_ptr->items[i]->value);
  printf("-------------------\n\n");
}

int main(void)
{
  HashMap *map_ptr = create_hash_map(MAX_CAPACITY);
  insert_into_hash_map(map_ptr, (char *)"1", (char *)"First address");
  insert_into_hash_map(map_ptr, (char *)"2", (char *)"Second address");
  insert_into_hash_map(map_ptr, (char *)"Hello", (char *)"Third address");
  insert_into_hash_map(map_ptr, (char *)"Awesome", (char *)"Fourth address");
  print_search_operation(map_ptr, (char *)"1");
  print_search_operation(map_ptr, (char *)"2");
  print_search_operation(map_ptr, (char *)"3");
  print_search_operation(map_ptr, (char *)"Hello");
  print_search_operation(map_ptr, (char *)"Sebastian");
  print_hash_map(map_ptr);
  delete_item_from_hash_map(map_ptr, (char *)"1");
  delete_item_from_hash_map(map_ptr, (char *)"LingLing");
  print_hash_map(map_ptr);
  free_hash_map(map_ptr);

  return 0;
}
