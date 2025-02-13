import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:intl/intl.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/core/providers/connectivity_provider.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/core/utils/tools.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/platform/mobile/widgets/chat.dart';
import 'package:squad_go/platform/mobile/widgets/custom_label.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/edit_event.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/map_location.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/offline.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/share_event.dart';
import 'package:squad_go/platform/mobile/widgets/score.dart';
import 'package:squad_go/platform/mobile/widgets/teams.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class EventScreen extends StatefulWidget {
  final Event? event;
  final String? eventId;

  const EventScreen({super.key, this.event, this.eventId})
      : assert(event != null || eventId != null);

  @override
  State<EventScreen> createState() => _EventScreenState();
}

class _EventScreenState extends State<EventScreen>
    with TickerProviderStateMixin {
  final EventService eventService = EventService();
  late Event event = widget.event ?? Event.empty();
  bool isOrganizer = false;
  bool isCoach = false;

  final DateTime today =
      DateTime.parse(DateFormat('yyyy-MM-dd').format(DateTime.now()));

  DateTime get eventDate => DateTime.parse(
          DateFormat('yyyy-MM-dd HH:mm').format(DateTime.parse(event.date)))
      .add(Duration(hours: 1));

  bool get isEventFinished => eventDate.isBefore(today);

  bool get isEventToday =>
      (eventDate.year == today.year &&
          eventDate.month == today.month &&
          eventDate.day == today.day) &&
      eventDate.isBefore(DateTime.now());

  late TabController _tabController;

  @override
  void initState() {
    super.initState();
    _tabController = TabController(length: 3, vsync: this);
    _tabController.addListener(() {
      setState(() {}); // Force le rebuild quand l'onglet change
    });
    _fetchEventDetails();
  }

  Future<void> _fetchEventDetails() async {
    try {
      final String eventId = widget.event?.id ?? widget.eventId!;
      final eventDetails = await eventService.getEventById(eventId);

      final currentUserId = context.read<AuthState>().userInfo?.id;
      final isUserCoach =
          hasRole(PlayerRole.coach, currentUserId, eventDetails.teams);
      final isUserOrg =
          hasRole(PlayerRole.org, currentUserId, eventDetails.teams);

      setState(() {
        event = eventDetails;
        isOrganizer = event.createdBy == currentUserId || isUserOrg;
        isCoach = isUserCoach;
      });
    } catch (e) {
      log.severe('Failed to fetch event details: $e');
    }
  }

  Future<void> onRefresh() async {
    await _fetchEventDetails();
  }

  @override
  Widget build(BuildContext context) {
    var isOnline = context.watch<ConnectivityState>().isConnected;

    final translate = AppLocalizations.of(context);
    return SafeArea(
      child: Scaffold(
        floatingActionButton: !isEventFinished && _tabController.index == 0
            ? FloatingActionButton(
                shape: const CircleBorder(),
                backgroundColor: event.sport.color?.withValues(alpha: 0.5) ??
                    Theme.of(context)
                        .colorScheme
                        .primary
                        .withValues(alpha: 0.5),
                onPressed: () {
                  showDialog(
                    context: context,
                    builder: (context) {
                      return ShareEventDialog(event: event);
                    },
                  );
                },
                child: const Icon(
                  Icons.share,
                  color: Colors.white,
                ),
              )
            : null,
        body: event.id == null
            ? const Center(child: CircularProgressIndicator())
            : RefreshIndicator(
                edgeOffset: 40,
                onRefresh: onRefresh,
                child: CustomScrollView(
                  physics: const AlwaysScrollableScrollPhysics(),
                  slivers: [
                    SliverAppBar(
                      leading: IconButton(
                        icon: const Icon(
                          Icons.arrow_back,
                          color: Colors.white,
                        ),
                        onPressed: () {
                          context.go('/home');
                        },
                      ),
                      actions: [
                        if (isOrganizer && !isEventFinished)
                          IconButton(
                            icon: const Icon(Icons.edit, color: Colors.white),
                            onPressed: () {
                              if (!isOnline) {
                                showDialog(
                                  context: context,
                                  builder: (context) => const OfflineDialog(),
                                );
                                return;
                              }
                              showDialog(
                                context: context,
                                builder: (context) {
                                  return EditEventDialog(
                                    event: event,
                                    onRefresh: onRefresh,
                                  );
                                },
                              );
                            },
                          ),
                      ],
                      title: Tooltip(
                        message: event.name,
                        child: Text(
                          event.name,
                          style: TextStyle(
                              fontWeight: FontWeight.bold, color: Colors.white),
                        ),
                      ),
                      pinned: true,
                      expandedHeight: event.sport.imageUrl != null ? 100 : 0,
                      flexibleSpace: FlexibleSpaceBar(
                        background: event.sport.imageUrl != null
                            ? Stack(
                                fit: StackFit.expand,
                                children: [
                                  Image.network(
                                    event.sport.imageUrl as String,
                                    fit: BoxFit.cover,
                                  ),
                                  Container(
                                    color: Colors.black
                                        .withAlpha(100), // Assombrit l'image
                                  ),
                                ],
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
                                bottom: 16,
                                left: 16,
                                right: 10,
                              ),
                              padding: const EdgeInsets.all(16),
                              decoration: BoxDecoration(
                                color: event.sport.color
                                        ?.withValues(alpha: 0.03) ??
                                    Theme.of(context)
                                        .colorScheme
                                        .primary
                                        .withValues(alpha: 0.03),
                                borderRadius: BorderRadius.all(
                                  Radius.circular(16),
                                ),
                              ),
                              child: Column(
                                mainAxisAlignment:
                                    MainAxisAlignment.spaceBetween,
                                children: [
                                  InkWell(
                                    onTap: () {
                                      showDialog(
                                        context: context,
                                        builder: (context) {
                                          return MapLocation(
                                            latitude: event.latitude,
                                            longitude: event.longitude,
                                          );
                                        },
                                      );
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
                                  Wrap(
                                    spacing: 8.0,
                                    runSpacing: 8.0,
                                    children: [
                                      CustomLabel(
                                        label: event.sport.name.name[0]
                                                .toUpperCase() +
                                            event.sport.name.name.substring(1),
                                        icon: sportIcon[event.sport.name],
                                        color: event.sport.color,
                                        iconColor: event.sport.color,
                                        backgroundColor:
                                            event.sport.color?.withAlpha(20),
                                      ),
                                      CustomLabel(
                                        label:
                                            event.type.name[0].toUpperCase() +
                                                event.type.name.substring(1),
                                        icon: eventTypeIcon[event.type],
                                        color: eventTypeColor[event.type],
                                        iconColor: eventTypeColor[event.type],
                                        backgroundColor:
                                            eventTypeColor[event.type]
                                                ?.withAlpha(20),
                                      ),
                                      CustomLabel(
                                        label: DateFormat('yyyy/MM/dd HH:mm')
                                            .format(
                                          DateTime.parse(event.date).add(
                                            Duration(hours: 1),
                                          ),
                                        ),
                                        icon: Icons.date_range,
                                        color: getColorBasedOnDate(event.date),
                                        iconColor:
                                            getColorBasedOnDate(event.date),
                                        backgroundColor:
                                            getColorBasedOnDate(event.date)
                                                .withAlpha(20),
                                      ),
                                    ],
                                  ),
                                ],
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
                                length: 3,
                                child: Column(
                                  children: [
                                    Container(
                                      height: 40,
                                      decoration: BoxDecoration(
                                        color: event.sport.color
                                                ?.withValues(alpha: 0.2) ??
                                            Theme.of(context)
                                                .colorScheme
                                                .primary
                                                .withValues(alpha: 0.2),
                                        borderRadius: BorderRadius.circular(10),
                                      ),
                                      child: TabBar(
                                        controller: _tabController,
                                        indicatorSize: TabBarIndicatorSize.tab,
                                        dividerColor: Colors.transparent,
                                        indicator: BoxDecoration(
                                          color: event.sport.color
                                                  ?.withValues(alpha: 0.5) ??
                                              Theme.of(context)
                                                  .colorScheme
                                                  .primary
                                                  .withValues(alpha: 0.5),
                                          borderRadius:
                                              BorderRadius.circular(10),
                                        ),
                                        labelColor: Colors.white,
                                        labelStyle: TextStyle(
                                          fontWeight: FontWeight.bold,
                                        ),
                                        tabs: [
                                          Tab(
                                            child: Text(
                                              translate?.teams ?? 'Équipes',
                                              overflow: TextOverflow.ellipsis,
                                            ),
                                          ),
                                          Tab(
                                            child: Text(
                                              'Chat',
                                              overflow: TextOverflow.ellipsis,
                                            ),
                                          ),
                                          Tab(
                                            child: Text(
                                              translate?.live_score ?? 'Live Score',
                                              overflow: TextOverflow.ellipsis,
                                            ),
                                          ),
                                        ],
                                      ),
                                    ),
                                    Expanded(
                                      child: TabBarView(
                                        controller: _tabController,
                                        children: [
                                          // Onglet Équipes
                                          Container(
                                            margin:
                                                const EdgeInsets.only(top: 16),
                                            decoration: BoxDecoration(
                                              color: event.sport.color
                                                      ?.withValues(
                                                          alpha: 0.03) ??
                                                  Theme.of(context)
                                                      .colorScheme
                                                      .primary
                                                      .withValues(alpha: 0.03),
                                              borderRadius:
                                                  BorderRadius.circular(16),
                                            ),
                                            padding: const EdgeInsets.all(16),
                                            child: event.id != null
                                                ? TeamsHandle(
                                                    eventId: event.id!,
                                                    maxTeams:
                                                        event.sport.maxTeams,
                                                    teams: event.teams ?? [],
                                                    userIsCoach: isCoach,
                                                    userIsOrganizer:
                                                        isOrganizer,
                                                    color: event.sport.color ??
                                                        Theme.of(context)
                                                            .colorScheme
                                                            .primary,
                                                    onRefresh: onRefresh,
                                                    eventCreatorId:
                                                        event.createdBy,
                                                    isEventFinished:
                                                        isEventFinished,
                                                    isEventNowPlaying:
                                                        isEventToday,
                                                  )
                                                : null,
                                          ),
                                          ChatPage(
                                            eventID: event.id ?? '',
                                            sportColor: event.sport.color,
                                            isEventFinished: isEventFinished,
                                          ),
                                          Container(
                                            child: event.id != null ?
                                            Score(
                                              teams: event.teams ?? [],
                                              eventId: event.id!,
                                              sportId: event.sport.id,
                                              isEventNowPlaying: isEventToday,
                                              isEventFinished: isEventFinished,
                                            ): null,
                                          ),
                                        ],
                                      ),
                                    )
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
      ),
    );
  }
}
