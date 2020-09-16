#include <stdio.h>
#include <stdlib.h>
#include "uthash.h"

typedef struct {
    int key;
    char in[20];
    int time;
    UT_hash_handle hh;
} Passenger;

typedef struct {
    char in[20];
    char out[20];
} StartEndHash;

typedef struct {
    StartEndHash key;
    int time[10000];
    int count;
    UT_hash_handle hh;
} Stations;

typedef struct {
    Passenger *pPassenger;
    Stations *pStations;
} UndergroundSystem;


UndergroundSystem* undergroundSystemCreate() {
    UndergroundSystem *pUS = malloc(sizeof(UndergroundSystem));
    pUS->pPassenger = NULL;
    pUS->pStations = NULL;
    return pUS;
}

void undergroundSystemCheckIn(UndergroundSystem* obj, int id, char * stationName, int t) {
    Passenger *p = NULL;
    HASH_FIND_INT(obj->pPassenger, &id, p);
    if (p == NULL) {
        p = malloc(sizeof(Passenger));
        p->key = id;
        HASH_ADD_INT(obj->pPassenger, key, p);
    }
    strcpy(p->in, stationName);
    p->time = t;
}

void undergroundSystemCheckOut(UndergroundSystem* obj, int id, char * stationName, int t) {
    Passenger *p = NULL;
    Stations *s = NULL;
    StartEndHash key;

    HASH_FIND_INT(obj->pPassenger, &id, p);
    if (p == NULL) {
        return;
    }
    
    memset(&key, 0, sizeof(key));
    strcpy(key.in, p->in);
    strcpy(key.out, stationName);

    HASH_FIND(hh, obj->pStations, &key, sizeof(key), s);
    if (s == NULL) {
        s = malloc(sizeof(Stations));
        memcpy(&s->key, &key, sizeof(key));
        s->count = 0;
        HASH_ADD(hh, obj->pStations, key, sizeof(key), s);
    }

    s->time[s->count++] = t - p->time;
}

double undergroundSystemGetAverageTime(UndergroundSystem* obj, char * startStation, char * endStation) {
    Stations *s = NULL;
    double sum = 0;
    StartEndHash key;

    memset(&key, 0, sizeof(key));
    strcpy(key.in, startStation);
    strcpy(key.out, endStation);
    HASH_FIND(hh, obj->pStations, &key, sizeof(key), s);
    if (s == NULL) {
        return 0.0;
    }
    for (int i = 0; i < s->count; i++) {
        sum += s->time[i];
    }
    return sum / s->count;
}

void undergroundSystemFree(UndergroundSystem* obj) {
    Stations *s = NULL, *ss;
    Passenger *p = NULL, *pp;
    HASH_ITER(hh, obj->pStations, s, ss) {
        HASH_DEL(obj->pStations, s);
        free(s);
    }
    HASH_ITER(hh, obj->pPassenger, p, pp) {
        HASH_DEL(obj->pPassenger, p);
        free(p);
    }
    free(obj);
}

/**
 * Your UndergroundSystem struct will be instantiated and called as such:
 * UndergroundSystem* obj = undergroundSystemCreate();
 * undergroundSystemCheckIn(obj, id, stationName, t);
 
 * undergroundSystemCheckOut(obj, id, stationName, t);
 
 * double param_3 = undergroundSystemGetAverageTime(obj, startStation, endStation);
 
 * undergroundSystemFree(obj);
*/
int main() {
    UndergroundSystem* pUS =  undergroundSystemCreate();
    undergroundSystemCheckIn(pUS, 45, "Leyton", 3);
    undergroundSystemCheckIn(pUS, 32, "Paradise", 8);
    undergroundSystemCheckIn(pUS, 27, "Leyton", 10);
    undergroundSystemCheckOut(pUS, 45, "Waterloo", 15);
    undergroundSystemCheckOut(pUS, 27, "Waterloo", 20);
    undergroundSystemCheckOut(pUS, 32, "Cambridge", 22);

    double param_3 = undergroundSystemGetAverageTime(pUS, "Leyton", "Waterloo");
    printf("Output: %f.", param_3);

    undergroundSystemFree(pUS);
    return 0;
}