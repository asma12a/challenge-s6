import 'dart:async';
import 'package:flutter/material.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/platform/mobile/widgets/event_card.dart';
import 'package:intl/intl.dart';

class SearchScreen extends StatefulWidget {
  const SearchScreen({super.key});

  @override
  State<SearchScreen> createState() => _SearchScreenState();
}

class _SearchScreenState extends State<SearchScreen> {
  List<Sport> _sports = [];
  bool _isLoading = true;
  List<Event> _searchResults = [];
  Map<String, String> params = {};
  Timer? _debounce;
  final eventService = EventService();

  @override
  void initState() {
    super.initState();
    _initSports();
    _getSearchResult();
  }

  Future<void> _initSports() async {
    try {
      final fetchedSports = await eventService.getSports();
      setState(() {
        _sports = fetchedSports;
      });
    } catch (error) {
      ScaffoldMessenger.of(context).clearSnackBars();
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text(
            textAlign: TextAlign.center,
            error.toString(),
            style: TextStyle(color: Theme.of(context).colorScheme.onError),
          ),
          backgroundColor: Theme.of(context).colorScheme.error,
        ),
      );
    }
  }

  _getSearchResult() async {
    try {
      final fetchedEvents = await eventService.getSearchResults(params);
      setState(() {
        _searchResults = fetchedEvents;
        debugPrint("Résultats de recherche : $_searchResults");
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
                params["search"] = value!;
                _getSearchResult();
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
                        value: sport.id,
                        child: Text(
                          sport.name.name,
                          style: TextStyle(
                              color: Theme.of(context).colorScheme.onSecondary),
                        ),
                      );
                    }).toList(),
                    onChanged: (String? value) {
                      params["sport"] = value!;
                      _getSearchResult();
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
            _searchResults.isNotEmpty
                ? "${_searchResults.length} événements trouvés."
                : "Aucun événement trouvé.",
          ),
          SizedBox(
            height: 5,
          ),
          Expanded(
            child: ListView.builder(
              itemCount: _searchResults.length,
              itemBuilder: (ctx, index) => EventCard(
               event: _searchResults[index],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
