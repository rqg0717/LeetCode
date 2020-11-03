#include <stdio.h>
#include <stdlib.h>

typedef struct NODE
{
   int value;
   struct NODE *front;
   struct NODE *back;
} node;

node *head = NULL;
node *tail = NULL;

void createMyQueue() {
    head = calloc(1, sizeof(node));
    tail = calloc(1, sizeof(node));
    head->value = 0;
    head->front = NULL;
    head->back = tail;
    tail->value = 0;
    tail->front = head;
    tail->back = NULL;
}

void enqueue(int value) {
    node *front = tail->front;
    node *tmp = calloc(1, sizeof(node));
    tmp->value = value;
    front->back = tmp;
    tail->front = tmp;
    tmp->front = front;
    tmp->back = tail;
}

void dequeue() {
    if (head->back == tail) {
        return;
    }
    node *back = head->back;
    head->back = head->back->back;
    free(back);
}

void deleteMyQueue() {
    node *front = head;
    while(front->back != NULL){
        node *back = front->back;
        free(front);
        front = back;
    }
}

int main() {
    createMyQueue();
    dequeue();
    enqueue(1);
    enqueue(2);
    enqueue(3);
    dequeue();
    dequeue();
    printf("myQueue Front is: %d \n", head->back->value);
    deleteMyQueue();
}