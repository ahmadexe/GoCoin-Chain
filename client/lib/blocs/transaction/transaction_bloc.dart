import 'package:client/models/transaction.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

part 'transaction_event.dart';
part 'transaction_state.dart';

part 'data_provider.dart';
part 'repository.dart';

class TransactionBloc extends Bloc<TransactionEvent, TransactionState> {
  static TransactionBloc get(context, [listen = true]) =>
      BlocProvider.of(context, listen: listen);
  TransactionBloc() : super(TransactionInitial()) {
    on<CreateTransaction>(_createTransaction);
  }

  Future<void> _createTransaction(
      CreateTransaction event, Emitter<TransactionState> emit) async {
    emit(
      TransactionLoading(),
    );

    try {
      await _TransactionRepo().createTransaction(event.transaction);

      emit(
        TransactionLoaded(),
      );
    } catch (e) {
      emit(
        const TransactionError(message: 'Failed to create transaction'),
      );
    }
  }
}
