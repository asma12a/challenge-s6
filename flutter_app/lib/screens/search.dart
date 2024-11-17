import 'dart:async';
import 'package:flutter/material.dart';
import 'package:flutter_app/core/services/event_service.dart';
import 'package:flutter_app/models/event.dart';
import 'package:flutter_app/models/sport.dart';
import 'package:flutter_app/widgets/event_card.dart';
import 'package:intl/intl.dart';

class SearchScreen extends StatefulWidget {
  const SearchScreen({super.key});

  @override
  State<SearchScreen> createState() => _SearchScreenState();
}

class _SearchScreenState extends State<SearchScreen> {
  List<Map<String, dynamic>> _sports = [];
  bool _isLoading = true;
  List<Map<String, dynamic>> _searchResults = [];
  Map<String, String> params = {};
  Timer? _debounce;

  @override
  void initState() {
    super.initState();
    _initSports();
    _getSearchResult();
  }

  Future<void> _initSports() async {
    try {
      final fetchedSports = await EventService.getSports();
      setState(() {
        _sports = fetchedSports;
        _isLoading = false;
      });
    } catch (error) {
      setState(() {
        _isLoading = false;
      });
    }
  }

  _getSearchResult() async {
    try {
      final fetchedEvents = await EventService.getSearchResults(params);
      setState(() {
        _searchResults = fetchedEvents;
        print(_searchResults);
        _isLoading = false;
      });
    } catch (error) {
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.all(15),
      child: Column(
        children: [
          SearchBar(
            leading: Icon(Icons.search),
            padding: WidgetStateProperty.all<EdgeInsets>(
              EdgeInsets.symmetric(horizontal: 20),
            ),
            onChanged: (String? value) {
              // Annuler le précédent Timer
              if (_debounce?.isActive ?? false) _debounce?.cancel();
              // Créer un nouveau Timer de 2 secondes
              _debounce = Timer(const Duration(seconds: 2), () {
                setState(() {
                  params["address"] = value!;
                  _getSearchResult();
                });
              });
            },
          ),
          SizedBox(
            height: 20,
          ),
          Row(
            children: [
              Expanded(
                child: Container(
                  padding: const EdgeInsets.all(5.0),
                  decoration: BoxDecoration(
                      color: Theme.of(context).colorScheme.secondary,
                      border: Border.all(color: Colors.grey),
                      borderRadius: BorderRadius.circular(8.0)),
                  child: DropdownButton<String>(
                    value: params["sport"],
                    iconEnabledColor: Colors.black,
                    dropdownColor: Theme.of(context).colorScheme.secondary,
                    hint: Text(
                      "Sport",
                      style: TextStyle(
                          color: Theme.of(context).colorScheme.onSecondary),
                    ),
                    items: _sports.map((sport) {
                      return DropdownMenuItem<String>(
                        value: sport['id'],
                        child: Text(
                          sport['name'],
                          style: TextStyle(
                              color: Theme.of(context).colorScheme.onSecondary),
                        ),
                      );
                    }).toList(),
                    onChanged: (String? value) {
                      setState(() {
                        params["sport"] = value!;
                        _getSearchResult();
                      });
                    },
                  ),
                ),
              ),
              SizedBox(
                width: 10,
              ),
              Expanded(
                child: Container(
                  padding: const EdgeInsets.all(5.0),
                  decoration: BoxDecoration(
                      color: Theme.of(context).colorScheme.secondary,
                      border: Border.all(color: Colors.grey),
                      borderRadius: BorderRadius.circular(8.0)),
                  child: DropdownButton<String>(
                    value: params["type"],
                    dropdownColor: Theme.of(context).colorScheme.secondary,
                    iconEnabledColor: Colors.black,
                    hint: Text(
                      "Type",
                      style: TextStyle(
                          color: Theme.of(context).colorScheme.onSecondary),
                    ),
                    underline: Container(),
                    items: [
                      DropdownMenuItem(
                        value: "match",
                        child: Text("Match"),
                      ),
                      DropdownMenuItem(
                          value: "training", child: Text("Training"))
                    ],
                    onChanged: (String? value) {
                      params["type"] = value!;
                      _getSearchResult();
                    },
                  ),
                ),
              ),
            ],
          ),
          SizedBox(
            height: 10,
          ),
          Text(
            _searchResults.length > 1
                ? "${_searchResults.length} résultats trouvés."
                : "${_searchResults.length} résultat trouvé.",
            style: TextStyle(color: Colors.white),
          ),
          SizedBox(
            height: 5,
          ),
          Expanded(
            child: ListView.builder(
              itemCount: _searchResults.length,
              itemBuilder: (ctx, index) => EventCard(
                event: Event(
                  id: _searchResults[index]["id"],
                  name: _searchResults[index]["name"],
                  date: DateFormat('yyyy-MM-dd').format(
                    DateTime.tryParse(_searchResults[index]["date"]) ??
                        DateTime.now(),
                  ),
                  address: _searchResults[index]["address"],
                  sport: Sport(
                      id: _searchResults[index]["sport"]["id"],
                      name: SportName.values.firstWhere(
                        (sn) => sn.name.contains(
                          _searchResults[index]["sport"]["name"]
                              .toString()
                              .toLowerCase(),
                        ),
                        orElse: () => SportName.football,
                      ),
                      type: SportType.values.firstWhere(
                        (st) => st.name.contains(
                          _searchResults[index]["sport"]["type"]
                              .toString()
                              .toLowerCase(),
                        ),
                        orElse: () => SportType.team,
                      ),
                      color: _searchResults[index]["sport"]["color"] != null
                          ? Color(int.parse(
                              _searchResults[index]["sport"]["color"]
                                  .toString(),
                              radix: 16,
                            ))
                          : Colors.black,
                      imageUrl: _searchResults[index]["sport"]["imageUrl"]),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
