import 'package:animated_toggle_switch/animated_toggle_switch.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/platform/mobile/widgets/carousel.dart';
import 'package:squad_go/platform/mobile/widgets/event_card.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class HomeMyEvents extends StatefulWidget {
  final Future<void> Function()? onRefresh;
  final bool? isHome;
  final Function(int)? onEventsCountChanged;
  final Function(List<Sport>)? onDistinctSportsFetched;

  const HomeMyEvents(
      {super.key,
      this.onRefresh,
      this.isHome,
      this.onEventsCountChanged,
      this.onDistinctSportsFetched});

  @override
  State<HomeMyEvents> createState() => HomeMyEventsState();
}

class HomeMyEventsState extends State<HomeMyEvents> {
  final EventService eventService = EventService();
  List<Event> allEvents = [];
  List<Event> myEvents = [];
  var _isPassed = true;

  @override
  void initState() {
    super.initState();

    fetchMyEvents();
  }

  Future<void> fetchMyEvents() async {
    try {
      List<Event> events = await eventService.getMyEvents();
      DateTime now = DateTime.now();
      final filteredEvents = events.where((event) {
        DateTime eventDate = DateTime.parse(event.date).toLocal();
        DateTime eventDay = DateTime(eventDate.year, eventDate.month,
            eventDate.day, eventDate.hour, eventDate.minute);
        return eventDate.year == now.year &&
                eventDate.month == now.month &&
                eventDate.day == now.day ||
            eventDay.isAfter(now);
      }).toList();

      setState(() {
        allEvents = events;
        myEvents = filteredEvents;
      });
      if (widget.onEventsCountChanged != null) {
        widget.onEventsCountChanged!(events.length);
      }
    } catch (e) {
      // Handle error
      log.severe('Failed to fetch events: $e');
    }
  }

  _filterEvents(value) {
    _isPassed = value;
    debugPrint(value.toString());
    DateTime now = DateTime.now();
    final DateTime today =
        DateTime.parse(DateFormat('yyyy-MM-dd').format(DateTime.now()));
    if (!value) {
      setState(() {
        myEvents = allEvents.where((event) {
          DateTime eventDate = DateTime.parse(event.date).toLocal();
          return eventDate.isBefore(today);
        }).toList();
      });
    } else {
      setState(() {
        myEvents = allEvents.where((event) {
          DateTime eventDate = DateTime.parse(event.date).toLocal();
          return eventDate.year == now.year &&
                  eventDate.month == now.month &&
                  eventDate.day == now.day ||
              eventDate.isAfter(today);
        }).toList();
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    return widget.isHome!
        ? Carousel(
            text: translate?.my_events ?? "Mes événements en cours",
            items: myEvents
                .map((event) => EventCard(
                      event: event,
                      hasJoinedEvent: true,
                      onRefresh: widget.onRefresh,
                    ))
                .toList(),
          )
        : Flexible(
            child: Column(
              children: [
                AnimatedToggleSwitch.dual(
                  first: false,
                  second: true,
                  current: _isPassed,
                  onChanged: (value) => _filterEvents(value),
                  iconBuilder: (value) => value
                      ? const Icon(
                          Icons.timer,
                          color: Colors.white,
                        )
                      : const Icon(
                          Icons.timer_off,
                          color: Colors.white,
                        ),
                  height: 40,
                  style: ToggleStyle(
                    indicatorColor: Colors.blue,
                    borderColor: Colors.blue,
                  ),
                  textBuilder: (value) => value
                      ? Text(
                          translate?.event_incoming ?? 'Actuel',
                          style: TextStyle(fontWeight: FontWeight.bold),
                        )
                      : Text(
                          translate?.event_past ?? 'Passés',
                          style: TextStyle(fontWeight: FontWeight.bold),
                        ),
                ),
                SizedBox(height: 20),
                myEvents.isEmpty
                    ? Center(
                        child: Text(translate?.no_event_to_display ??
                            'Aucun événement à afficher'),
                      )
                    : Expanded(
                        child: ListView.builder(
                          itemCount: myEvents.length,
                          itemBuilder: (ctx, index) => EventCard(
                            event: myEvents[index],
                            hasJoinedEvent: true,
                            onRefresh: widget.onRefresh,
                          ),
                        ),
                      ),
              ],
            ),
          );
  }
}
