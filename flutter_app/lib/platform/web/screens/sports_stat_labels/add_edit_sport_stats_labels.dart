import 'package:flutter/material.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';

class AddEditSportStatLabelModal extends StatefulWidget {
  final SportStatLabels? statLabel;
  final Function onStatLabelSaved;
  final List<Map<String, dynamic>> sports; // Liste des sports disponibles

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
      _unit = widget.statLabel!.unit;
      _isMain = widget.statLabel!.isMain;
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
            'sport_id': _selectedSportId, // Ajout ou mise à jour de sport_id
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
    return Dialog(
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12.0)),
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Text(
                widget.statLabel == null
                    ? 'Ajouter une Statistique'
                    : 'Modifier la Statistique',
                style: const TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 16),
              TextFormField(
                initialValue: _label,
                decoration: const InputDecoration(labelText: 'Nom'),
                validator: (value) =>
                    value == null || value.isEmpty ? 'Nom requis' : null,
                onSaved: (value) => _label = value!,
              ),
              const SizedBox(height: 16),
              TextFormField(
                initialValue: _unit,
                decoration: const InputDecoration(labelText: 'Unité'),
                validator: (value) =>
                    value == null || value.isEmpty ? 'Unité requise' : null,
                onSaved: (value) => _unit = value!,
              ),
              const SizedBox(height: 16),
              DropdownButtonFormField<String>(
                value: _selectedSportId,
                decoration: const InputDecoration(labelText: 'Sport'),
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
                validator: (value) =>
                    value == null ? 'Veuillez sélectionner un sport' : null,
              ),
              const SizedBox(height: 16),
              CheckboxListTile(
                title: const Text('Décisif'),
                value: _isMain,
                onChanged: (value) {
                  setState(() {
                    _isMain = value!;
                  });
                },
              ),
              const SizedBox(height: 16),
              ElevatedButton(
                onPressed: _isSaving ? null : _saveStatLabel,
                child: _isSaving
                    ? const CircularProgressIndicator()
                    : Text(widget.statLabel == null ? 'Créer' : 'Modifier'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
