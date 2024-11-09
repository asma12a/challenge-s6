import 'package:flutter/material.dart';
import 'package:flutter_app/widgets/event_search.dart';

class SearchScreen extends StatefulWidget {
  const SearchScreen({super.key});

  @override
  State<SearchScreen> createState() => _SearchScreenState();
}

class _SearchScreenState extends State<SearchScreen> {
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
                    iconDisabledColor: Colors.black,
                    hint: Text(
                      "Sport",
                      style: TextStyle(
                          color: Theme.of(context).colorScheme.onSecondary),
                    ),
                    underline: Container(),
                    items: [],
                    onChanged: (String? value) {},
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
                    iconDisabledColor: Colors.black,
                    hint: Text(
                      "Type",
                      style: TextStyle(
                          color: Theme.of(context).colorScheme.onSecondary),
                    ),
                    underline: Container(),
                    items: [],
                    onChanged: (String? value) {},
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
                    iconDisabledColor: Colors.black,
                    hint: Text(
                      "Statut",
                      style: TextStyle(
                          color: Theme.of(context).colorScheme.onSecondary),
                    ),
                    underline: Container(),
                    items: [],
                    onChanged: (String? value) {},
                  ),
                ),
              ),
            ],
          ),
          SizedBox(
            height: 50,
          ),
          Text(
            "10 résultats trouvés.",
            style: TextStyle(color: Colors.white),
          ),
          Expanded(
            child: ListView.builder(
              itemCount: 10,
              itemBuilder: (ctx, index) => EventSearch(),
            ),
          ),
        ],
      ),
    );
  }
}
