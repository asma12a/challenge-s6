import 'package:flutter/material.dart';
import 'package:flutter_app/models/event.dart';
import 'package:flutter_app/models/sport.dart';

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
                            color: event.sport.color ??
                                Theme.of(context)
                                    .colorScheme
                                    .secondary
                                    .withAlpha(30),
                          ),
                  ),
                  Expanded(
                    flex: 4,
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        ListTile(
                          leading: const Icon(Icons.sports),
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
                                    Text(
                                      event.address,
                                      overflow: TextOverflow.ellipsis,
                                    ),
                                  ],
                                ),
                              ),
                            ],
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.only(
                            left: 16,
                            right: 16,
                            bottom: 16,
                          ),
                          child: Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              Row(
                                children: [
                                  Icon(Icons.date_range),
                                  SizedBox(width: 8),
                                  Text(event.date),
                                ],
                              ),
                              Flexible(
                                child: TextButton(
                                  onPressed: onCardClick,
                                  child: Text(
                                    hasJoinedEvent ? "VOIR" : 'REJOINDRE',
                                  ),
                                ),
                              )
                            ],
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
                  sportIcon[event.sport.name],
                  size: 32,
                  color: event.sport.color ??
                      Theme.of(context).colorScheme.secondary.withAlpha(30),
                ),
              ),
            ]),
          ),
        ),
      ),
    );
  }
}
