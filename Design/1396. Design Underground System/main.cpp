#include <string>
#include <unordered_map>
#include <iostream>

using namespace std;

class UndergroundSystem
{
public:
    using station = pair<string, int>;
    using Out = pair<string, string>;
    using Sum = pair<int, int>;

    struct StartEndHash
    {
        int operator()(const Out &stations) const
        {
            return hash<string>()(stations.first + stations.second);
        }
    };

    unordered_map<int, station> In;
    unordered_map<Out, Sum, StartEndHash> table;

    UndergroundSystem() {}

    void checkIn(int id, string stationName, int t)
    {
        In[id] = {stationName, t};
    }

    void checkOut(int id, string stationName, int t)
    {
        auto time = In[id].second;
        table[{In[id].first, stationName}].first += t - time;
        table[{In[id].first, stationName}].second++; // count
    }

    double getAverageTime(string startStation, string endStation)
    {
        auto sum = table[{startStation, endStation}].first;
        auto count = table[{startStation, endStation}].second;
        return double(sum) / count;
    }
};

int main()
{
    UndergroundSystem *pUS = new UndergroundSystem();
    pUS->checkIn(45, "Leyton", 3);
    pUS->checkIn(32, "Paradise", 8);
    pUS->checkIn(27, "Leyton", 10);
    pUS->checkOut(45, "Waterloo", 15);
    pUS->checkOut(27, "Waterloo", 20);
    pUS->checkOut(32, "Cambridge", 22);
    double param_3 = pUS->getAverageTime("Leyton", "Waterloo");
    cout << param_3 << endl;
    return 0;
}