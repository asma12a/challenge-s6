import 'package:flutter/material.dart';
import 'package:flutter_app/models/event.dart';

class EventCard extends StatelessWidget {
  const EventCard({super.key, required this.event});

  final Event event;

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Card(
        surfaceTintColor: Theme.of(context).colorScheme.tertiary,
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: <Widget>[
            ListTile(
              leading: const Icon(Icons.sports),
              title: Text(event.name),
              subtitle: Row(
                children: [
                  Expanded(
                      child: Text(
                    event.address,
                    overflow: TextOverflow.ellipsis,
                  )),
                  SizedBox(width: 8),
                  Icon(
                    Icons.place,
                    size: 16,
                  ),
                ],
              ),
            ),
            Stack(
              children: [
                // Image.network(
                //     height: 300,
                //     fit: BoxFit.cover,
                //     width: double.infinity,
                //     "https://images.unsplash.com/photo-1729592088218-02a52acb3547?q=80&w=2876&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"),
                Positioned(
                  bottom: 0,
                  left: 0,
                  right: 0,
                  child: Container(
                    color: Colors.black54,
                    padding:
                        const EdgeInsets.symmetric(vertical: 6, horizontal: 24),
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.start,
                      children: [
                        Icon(Icons.date_range),
                        SizedBox(
                          width: 8,
                        ),
                        Text(
                          event.date,
                          style: TextStyle(color: Colors.white),
                        ),
                        SizedBox(
                          width: 30,
                        ),
                        Icon(Icons.groups),
                        SizedBox(
                          width: 8,
                        ),
                        Text(
                          event.sport,
                          style: TextStyle(color: Colors.white),
                        ),
                      ],
                    ),
                  ),
                ),
              ],
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.end,
              children: <Widget>[
                TextButton(
                  child: const Text('REJOINDRE'),
                  onPressed: () {
                    /* ... */
                  },
                ),
                const SizedBox(width: 8),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
