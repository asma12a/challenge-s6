import 'package:flutter/material.dart';

class EventSearch extends StatelessWidget {
  const EventSearch({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 140,
      margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8.0),
      decoration: BoxDecoration(
        color: Theme.of(context).colorScheme.onTertiaryFixed,
        border: Border.all(color: const Color(0xFFE0E0E0)),
        borderRadius: BorderRadius.circular(8.0),
      ),
      padding: const EdgeInsets.all(8),
      child: Row(
        children: [
          Expanded(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  "Événement 1",
                  style: const TextStyle(
                      fontWeight: FontWeight.bold, color: Colors.white),
                  maxLines: 2,
                  overflow: TextOverflow.ellipsis,
                ),
                const SizedBox(height: 8),
                Text(
                  "Bass . Yesterday",
                  style: const TextStyle(
                      fontWeight: FontWeight.bold, color: Colors.white),
                ),
                const SizedBox(height: 8),
                Row(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Icons.bookmark_border_rounded,
                    Icons.share,
                    Icons.more_vert,
                  ].map((icon) {
                    return InkWell(
                      onTap: () {},
                      child: Padding(
                        padding: const EdgeInsets.only(right: 8.0),
                        child: Icon(icon, size: 16),
                      ),
                    );
                  }).toList(),
                ),
              ],
            ),
          ),
          Container(
            width: 100,
            height: 100,
            decoration: BoxDecoration(
              color: Colors.grey,
              borderRadius: BorderRadius.circular(8.0),
              image: DecorationImage(
                fit: BoxFit.cover,
                image: NetworkImage(
                  "https://images.unsplash.com/photo-1729592088218-02a52acb3547?q=80&w=2876&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
                ),
              ),
            ),
          )
        ],
      ),
    );
  }
}
