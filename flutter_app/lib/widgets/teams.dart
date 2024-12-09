import 'package:flutter/material.dart';
import 'package:squad_go/models/team.dart';

class TeamsHandle extends StatefulWidget {
  final String eventId;
  final List<Team> teams;
  final int? maxTeams;

  const TeamsHandle({
    super.key,
    required this.teams,
    required this.eventId,
    this.maxTeams,
  });

  @override
  State<TeamsHandle> createState() => _TeamsHandleState();
}

class _TeamsHandleState extends State<TeamsHandle> {
  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: widget.teams.length,
      child: Column(
        children: [
          Stack(
            children: [
              SizedBox(
                height: 40,
                child: TabBar(
                  isScrollable: true,
                  padding: EdgeInsets.only(right: 100),
                  dividerColor: Colors.transparent,
                  indicatorSize: TabBarIndicatorSize.tab,
                  tabAlignment: TabAlignment.start,
                  indicator: BoxDecoration(
                    color: Theme.of(context)
                        .colorScheme
                        .secondary
                        .withOpacity(0.5),
                    borderRadius: BorderRadius.circular(10),
                  ),
                  labelColor: Colors.white,
                  labelStyle: TextStyle(
                    fontWeight: FontWeight.bold,
                  ),
                  labelPadding: EdgeInsets.only(
                    left: 16,
                  ),
                  tabs: widget.teams.map(
                    (team) {
                      return Tab(
                        child: Row(
                          children: [
                            Text(
                              team.name,
                            ),
                            IconButton(
                              icon: Icon(Icons.edit),
                              onPressed: () {
                                // Show dialog to edit team name
                              },
                            ),
                          ],
                        ),
                      );
                    },
                  ).toList(),
                ),
              ),
              if (widget.maxTeams == null ||
                  (widget.maxTeams != null &&
                      widget.teams.length < widget.maxTeams!))
                Positioned(
                  right: 0,
                  child: SizedBox(
                    height: 40,
                    child: ElevatedButton(
                      onPressed: () {},
                      style: ElevatedButton.styleFrom(
                        shape: CircleBorder(),
                      ),
                      child: Icon(Icons.add),
                    ),
                  ),
                )
            ],
          ),
          Expanded(
            child: TabBarView(
              children: widget.teams
                  .map(
                    (team) => Container(
                      margin: EdgeInsets.only(top: 16),
                      child: Column(
                        children: [
                          Row(
                            children: [
                              Badge(
                                label: Text(
                                    '${team.players.length}/${team.maxPlayers}'),
                                backgroundColor:
                                    team.players.length < team.maxPlayers
                                        ? Colors.green.shade300
                                        : Colors.red.shade300,
                                textStyle: TextStyle(
                                  fontWeight: FontWeight.bold,
                                ),
                                padding: EdgeInsets.symmetric(
                                  horizontal: 7,
                                  vertical: 5,
                                ),
                              )
                            ],
                          ),
                          SizedBox(height: 16),
                          team.players.isNotEmpty
                              ? Expanded(
                                  child: ListView.builder(
                                    itemCount: team.players.length,
                                    itemBuilder: (context, index) {
                                      final player = team.players[index];
                                      return GestureDetector(
                                        onTap: () {
                                          // Show player details in dialog
                                        },
                                        child: Container(
                                          decoration: BoxDecoration(
                                            border: Border.all(
                                              color: Colors.grey,
                                              width: 1,
                                            ),
                                            borderRadius:
                                                BorderRadius.circular(16),
                                          ),
                                          child: ListTile(
                                            title: Text(player.name),
                                            trailing: Row(
                                              mainAxisSize: MainAxisSize.min,
                                              children: [
                                                IconButton(
                                                  icon: Icon(Icons.edit),
                                                  onPressed: () {
                                                    // Edit player details
                                                  },
                                                ),
                                                IconButton(
                                                  icon: Icon(Icons.bar_chart),
                                                  onPressed: () {
                                                    // Show player stats
                                                  },
                                                ),
                                              ],
                                            ),
                                          ),
                                        ),
                                      );
                                    },
                                  ),
                                )
                              : SizedBox(),
                          if (team.players.length < team.maxPlayers)
                            ElevatedButton(
                              onPressed: () {
                                // Show dialog to add player to team
                              },
                              child: Text('Add Player'),
                            ),
                        ],
                      ),
                    ),
                  )
                  .toList(),
            ),
          ),
        ],
      ),
    );
  }
}
