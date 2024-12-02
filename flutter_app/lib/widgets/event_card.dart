import 'package:flutter/material.dart';
import 'package:squad_go/models/event.dart';
import 'package:squad_go/models/sport.dart';
import 'package:intl/intl.dart';

class EventCard extends StatelessWidget {
  final Event event;
  final bool hasJoinedEvent;

  const EventCard(
      {super.key, required this.event, this.hasJoinedEvent = false});

  // make the on card click
  void onCardClick() {
    debugPrint('Card tapped.');
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Card(
        shadowColor: event.sport.color,
        clipBehavior: Clip.hardEdge,
        child: InkWell(
          splashColor: Theme.of(context).colorScheme.secondary.withAlpha(30),
          onTap: onCardClick,
          child: SizedBox(
            width: 300,
            height: 200,
            child: Stack(children: [
              Column(
                children: [
                  Expanded(
                    flex: 3,
                    child: event.sport.imageUrl != null
                        ? Image.network(
                            fit: BoxFit.cover,
                            width: double.infinity,
                            event.sport.imageUrl as String,
                          )
                        : Icon(
                            Icons.hide_image,
                            size: 80,
                            color: event.sport.color?.withAlpha(50) ??
                                Theme.of(context)
                                    .colorScheme
                                    .secondary
                                    .withAlpha(50),
                          ),
                  ),
                  Expanded(
                    flex: 4,
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        ListTile(
                          leading: Icon(
                            sportIcon[event.sport.name],
                            size: 32,
                            color: event.sport.color,
                          ),
                          title: Text(event.name),
                          subtitle: Row(
                            children: [
                              Expanded(
                                child: Row(
                                  children: [
                                    Icon(
                                      Icons.place,
                                      size: 16,
                                    ),
                                    SizedBox(width: 8),
                                    Expanded(
                                      child: Text(
                                        event.address,
                                        overflow: TextOverflow.ellipsis,
                                      ),
                                    ),
                                  ],
                                ),
                              ),
                            ],
                          ),
                        ),
                        Expanded(
                          child: Padding(
                            padding: const EdgeInsets.only(
                              left: 16,
                              right: 16,
                            ),
                            child: Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                Container(
                                  decoration: BoxDecoration(
                                    color: Theme.of(context)
                                        .colorScheme
                                        .secondary
                                        .withAlpha(30),
                                    borderRadius: BorderRadius.circular(8),
                                  ),
                                  padding: EdgeInsets.all(3),
                                  child: Row(
                                    children: [
                                      Icon(Icons.date_range),
                                      SizedBox(width: 8),
                                      Text(DateFormat('dd/MM/yyyy')
                                          .format(DateTime.parse(event.date))),
                                    ],
                                  ),
                                ),
                                TextButton(
                                  onPressed: onCardClick,
                                  child: Text(
                                    hasJoinedEvent ? "VOIR" : 'REJOINDRE',
                                  ),
                                )
                              ],
                            ),
                          ),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
              Positioned(
                right: 10,
                top: 10,
                child: Icon(
                  event.type == EventType.match
                      ? Icons.sports
                      : Icons.fitness_center,
                  size: 32,
                ),
              ),
            ]),
          ),
        ),
      ),
    );
  }
}
