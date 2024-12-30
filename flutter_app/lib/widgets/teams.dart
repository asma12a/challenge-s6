import 'package:flutter/material.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:provider/provider.dart';

class TeamsHandle extends StatefulWidget {
  final String eventId;
  final List<Team> teams;
  final int maxTeams;
  final bool canEdit;
  final Color color;

  const TeamsHandle({
    super.key,
    required this.teams,
    required this.eventId,
    required this.canEdit,
    required this.maxTeams,
    required this.color,
  });

  @override
  State<TeamsHandle> createState() => _TeamsHandleState();
}

class _TeamsHandleState extends State<TeamsHandle> {
  UserApp? currentUser;
  bool userHasTeam = false;

  @override
  void initState() {
    super.initState();

    currentUser = context.read<AuthState>().userInfo;
    userHasTeam = widget.teams.any(
        (team) => team.players.any((player) => player.id == currentUser!.id));
  }

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
                            SizedBox(width: 8),
                            if (widget.canEdit) ...[
                              Container(
                                width: 30,
                                height: 30,
                                decoration: BoxDecoration(
                                  shape: BoxShape.circle,
                                  color: Theme.of(context)
                                      .colorScheme
                                      .secondary
                                      .withOpacity(0.1),
                                ),
                                child: IconButton(
                                  icon: Icon(
                                    Icons.edit,
                                    size: 16,
                                  ),
                                  padding: EdgeInsets.zero,
                                  onPressed: () {
                                    // TODO: Edit team Dialog
                                  },
                                ),
                              ),
                              SizedBox(width: 8),
                            ],
                          ],
                        ),
                      );
                    },
                  ).toList(),
                ),
              ),
              if (widget.canEdit &&
                  (widget.maxTeams == 0 ||
                      widget.teams.length < widget.maxTeams))
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
                      margin: EdgeInsets.only(top: 8),
                      child: Column(
                        children: [
                          Row(
                            children: [
                              Badge(
                                label: Text(
                                    '${team.players.length}${team.maxPlayers > 0 ? '/${team.maxPlayers}' : ''}'),
                                backgroundColor: team.maxPlayers == 0
                                    ? Colors.grey.shade400
                                    : team.players.length < team.maxPlayers
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
                                            color: widget.color.withAlpha(20),
                                            borderRadius:
                                                BorderRadius.circular(16),
                                          ),
                                          child: ListTile(
                                            title: Text(
                                              player.name ?? player.email,
                                            ),
                                            trailing: Row(
                                              mainAxisSize: MainAxisSize.min,
                                              children: [
                                                if (widget.canEdit) ...[
                                                  IconButton(
                                                    icon: Icon(Icons.edit),
                                                    onPressed: () {
                                                      // TODO: Edit player Dialog
                                                    },
                                                  ),
                                                  IconButton(
                                                    icon: Icon(Icons.bar_chart),
                                                    onPressed: () {
                                                      // TODO: Show/Edit player stats Dialog
                                                    },
                                                  ),
                                                ],
                                                if (!widget.canEdit)
                                                  IconButton(
                                                    icon: Icon(
                                                      Icons.remove_red_eye,
                                                    ),
                                                    onPressed: () {
                                                      // TODO: Show player details Dialog
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
                          if (widget.canEdit &&
                              (team.maxPlayers == 0 ||
                                  team.players.length < team.maxPlayers))
                            ElevatedButton(
                              onPressed: () {
                                // TODO: Show dialog to add player to team
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
