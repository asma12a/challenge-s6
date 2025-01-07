import 'dart:ui';
import 'package:flutter/material.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';

class AddEditSportStatLabelModal extends StatefulWidget {
  final SportStatLabels? statLabel;
  final Function onStatLabelSaved;
  final List<Map<String, dynamic>> sports;

  const AddEditSportStatLabelModal({
    super.key,
    this.statLabel,
    required this.onStatLabelSaved,
    required this.sports,
  });

  @override
  State<AddEditSportStatLabelModal> createState() =>
      _AddEditSportStatLabelModalState();
}

class _AddEditSportStatLabelModalState
    extends State<AddEditSportStatLabelModal> {
  final _formKey = GlobalKey<FormState>();
  final SportStatLabelsService _statLabelsService = SportStatLabelsService();

  String _label = '';
  String _unit = '';
  bool _isMain = false;
  String? _selectedSportId;
  bool _isSaving = false;

  @override
  void initState() {
    super.initState();
    if (widget.statLabel != null) {
      _label = widget.statLabel!.label;
      _unit = widget.statLabel!.unit!;
      _isMain = widget.statLabel!.isMain!;
    }
  }

  Future<void> _saveStatLabel() async {
    if (!_formKey.currentState!.validate() || _selectedSportId == null) {
      if (_selectedSportId == null) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Veuillez sélectionner un sport.')),
        );
      }
      return;
    }

    _formKey.currentState!.save();
    setState(() {
      _isSaving = true;
    });

    try {
      final statLabelData = {
        'label': _label,
        'unit': _unit,
        'is_main': _isMain,
        'sport_id': _selectedSportId,
      };

      if (widget.statLabel == null) {
        await _statLabelsService.createStatLabel(statLabelData);
      } else {
        await _statLabelsService.updateStatLabel(
          {
            ...statLabelData,
            'sport_id': _selectedSportId,
          },
          widget.statLabel!.id.toString(),
        );
      }

      widget.onStatLabelSaved();
      Navigator.of(context).pop();
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur : $e')),
      );
    } finally {
      setState(() {
        _isSaving = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return BackdropFilter(
      filter: ImageFilter.blur(
          sigmaX: 5.0, sigmaY: 5.0), // Applique un flou au fond
      child: Dialog(
        shape:
            RoundedRectangleBorder(borderRadius: BorderRadius.circular(16.0)),
        backgroundColor: Colors.white,
        child: LayoutBuilder(
          builder: (context, constraints) {
            return Container(
              width: constraints.maxWidth *
                  0.6, // Réduction supplémentaire à 60% de largeur
              padding:
                  const EdgeInsets.symmetric(vertical: 24.0, horizontal: 20.0),
              child: Form(
                key: _formKey,
                child: SingleChildScrollView(
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Center(
                        child: Text(
                          widget.statLabel == null
                              ? 'Ajouter une Statistique'
                              : 'Modifier la Statistique',
                          style:
                              Theme.of(context).textTheme.titleLarge?.copyWith(
                                    fontWeight: FontWeight.bold,
                                    color: Colors.black87,
                                  ),
                        ),
                      ),
                      const SizedBox(height: 20),
                      TextFormField(
                        initialValue: _label,
                        decoration: InputDecoration(
                          labelText: 'Nom',
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8.0),
                          ),
                        ),
                        validator: (value) => value == null || value.isEmpty
                            ? 'Nom requis'
                            : null,
                        onSaved: (value) => _label = value!,
                      ),
                      const SizedBox(height: 16),
                      TextFormField(
                        initialValue: _unit,
                        decoration: InputDecoration(
                          labelText: 'Unité',
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8.0),
                          ),
                        ),
                        validator: (value) => value == null || value.isEmpty
                            ? 'Unité requise'
                            : null,
                        onSaved: (value) => _unit = value!,
                      ),
                      const SizedBox(height: 16),
                      DropdownButtonFormField<String>(
                        value: _selectedSportId,
                        decoration: InputDecoration(
                          labelText: 'Sport',
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8.0),
                          ),
                        ),
                        items: widget.sports.map((sport) {
                          return DropdownMenuItem<String>(
                            value: sport['id'].toString(),
                            child: Text(sport['name']),
                          );
                        }).toList(),
                        onChanged: (value) {
                          setState(() {
                            _selectedSportId = value;
                          });
                        },
                        validator: (value) => value == null
                            ? 'Veuillez sélectionner un sport'
                            : null,
                      ),
                      const SizedBox(height: 16),
                      CheckboxListTile(
                        contentPadding: EdgeInsets.zero,
                        title: const Text(
                          'Décisif',
                          style: TextStyle(fontSize: 14.0),
                        ),
                        activeColor: Theme.of(context).primaryColor,
                        value: _isMain,
                        onChanged: (value) {
                          setState(() {
                            _isMain = value!;
                          });
                        },
                      ),
                      const SizedBox(height: 16),
                      SizedBox(
                        width: double.infinity,
                        child: ElevatedButton(
                          style: ElevatedButton.styleFrom(
                            backgroundColor: Theme.of(context).primaryColor,
                            padding: const EdgeInsets.symmetric(vertical: 12.0),
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(8.0),
                            ),
                          ),
                          onPressed: _isSaving ? null : _saveStatLabel,
                          child: _isSaving
                              ? const CircularProgressIndicator(
                                  valueColor: AlwaysStoppedAnimation<Color>(
                                      Colors.white),
                                )
                              : Text(
                                  widget.statLabel == null
                                      ? 'Créer'
                                      : 'Modifier',
                                  style: const TextStyle(
                                    fontSize: 16.0,
                                    color: Colors.white,
                                  ),
                                ),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            );
          },
        ),
      ),
    );
  }
}
