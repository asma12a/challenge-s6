import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/widgets/custom_label.dart';
import 'package:squad_go/widgets/teams.dart';

class EventScreen extends StatelessWidget {
  final Event event;
  const EventScreen({super.key, required this.event});

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        body: CustomScrollView(
          slivers: [
            SliverAppBar(
              actions: [
                // TODO: DIsplay only if the user is the event creator
                IconButton(
                  icon: const Icon(Icons.edit),
                  onPressed: () {
                    // Edit event
                  },
                ),
              ],
              title: Tooltip(
                message: event.name,
                child: Text(
                  event.name,
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
              // floating: true,
              // snap: true,

              pinned: true,
              expandedHeight: event.sport.imageUrl != null ? 100 : 0,
              flexibleSpace: FlexibleSpaceBar(
                background: event.sport.imageUrl != null
                    ? Image.network(
                        event.sport.imageUrl as String,
                        fit: BoxFit.cover,
                      )
                    : null,
              ),
            ),
            SliverFillRemaining(
              child: Column(
                children: [
                  Expanded(
                    flex: 1,
                    child: Container(
                      margin: const EdgeInsets.only(
                          bottom: 16, left: 16, right: 16),
                      decoration: BoxDecoration(
                        color: Theme.of(context)
                            .colorScheme
                            .primary
                            .withOpacity(0.05),
                        borderRadius: BorderRadius.only(
                          bottomLeft: Radius.circular(16),
                          bottomRight: Radius.circular(16),
                        ),
                      ),
                      child: Padding(
                        padding: const EdgeInsets.all(16),
                        child: Column(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: [
                            InkWell(
                              onTap: () {
                                // Open map modal with the event location
                              },
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
                            Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                CustomLabel(
                                  label:
                                      event.sport.name.name[0].toUpperCase() +
                                          event.sport.name.name.substring(1),
                                  icon: sportIcon[event.sport.name],
                                  color: event.sport.color,
                                  iconColor: event.sport.color,
                                  backgroundColor:
                                      event.sport.color?.withAlpha(20),
                                ),
                                CustomLabel(
                                  label: event.type.name[0].toUpperCase() +
                                      event.type.name.substring(1),
                                  icon: eventTypeIcon[event.type],
                                  color: eventTypeColor[event.type],
                                  iconColor: eventTypeColor[event.type],
                                  backgroundColor:
                                      eventTypeColor[event.type]?.withAlpha(20),
                                ),
                                CustomLabel(
                                  label: DateFormat('dd/MM/yyyy')
                                      .format(DateTime.parse(event.date)),
                                  icon: Icons.date_range,
                                  color: Colors
                                      .green, // TODO: Change color based on date
                                  iconColor: Colors.green,
                                  backgroundColor: Colors.green.withAlpha(20),
                                ),
                              ],
                            ),
                          ],
                        ),
                      ),
                    ),
                  ),
                  Flexible(
                    flex: 4,
                    child: Container(
                      margin: const EdgeInsets.only(
                        bottom: 16,
                        left: 16,
                        right: 16,
                      ),
                      child: DefaultTabController(
                        length: 2,
                        child: Column(
                          children: [
                            Container(
                              height: 40,
                              decoration: BoxDecoration(
                                color: Theme.of(context)
                                    .colorScheme
                                    .primary
                                    .withOpacity(0.2),
                                borderRadius: BorderRadius.circular(10),
                              ),
                              child: TabBar(
                                indicatorSize: TabBarIndicatorSize.tab,
                                dividerColor: Colors.transparent,
                                indicator: BoxDecoration(
                                  color: Theme.of(context)
                                      .colorScheme
                                      .primary
                                      .withOpacity(0.5),
                                  borderRadius: BorderRadius.circular(10),
                                ),
                                labelColor: Colors.white,
                                labelStyle: TextStyle(
                                  fontWeight: FontWeight.bold,
                                ),
                                tabs: [
                                  Tab(
                                    child: Text(
                                      'Équipes',
                                      overflow: TextOverflow.ellipsis,
                                    ),
                                  ),
                                  Tab(
                                    child: Text(
                                      'Chat',
                                      overflow: TextOverflow.ellipsis,
                                    ),
                                  ),
                                ],
                              ),
                            ),
                            Expanded(
                              child: TabBarView(
                                children: [
                                  Container(
                                    margin: const EdgeInsets.only(top: 16),
                                    decoration: BoxDecoration(
                                      color: Theme.of(context)
                                          .colorScheme
                                          .primary
                                          .withOpacity(0.05),
                                      borderRadius: BorderRadius.circular(16),
                                    ),
                                    padding: const EdgeInsets.all(16),
                                    child: event.id != null
                                        ? TeamsHandle(
                                            eventId: event.id!,
                                            maxTeams: event.sport.maxTeams,
                                            teams: [
                                              Team(
                                                  id: "dsqdsq",
                                                  name: "Equipe A",
                                                  maxPlayers: 11,
                                                  players: [
                                                    Player(
                                                      id: "dsqdsq",
                                                      name: "Joueur 1",
                                                      email: "user@test.com",
                                                      role: PlayerRole.player,
                                                      status:
                                                          PlayerStatus.valid,
                                                    )
                                                  ]),
                                              Team(
                                                id: "dsqdsq",
                                                name: "Equipe B",
                                                maxPlayers: 11,
                                              ),
                                            ],
                                          )
                                        : Container(),
                                  ),
                                  Container(
                                    margin: const EdgeInsets.only(top: 16),
                                    color: Colors.blue,
                                    child: Center(child: Text('Chat')),
                                  ),
                                ],
                              ),
                            ),
                          ],
                        ),
                      ),
                    ),
                  )
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
