import { CommonModule } from '@angular/common';
import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { MatrixResponse } from '../../core/domain/dto/matrix/matrix.dto';
import { MatrixService } from '../../core/services/matrix/matrix.service';

@Component({
  selector: 'app-matrix.component',
  imports: [FormsModule, CommonModule],
  templateUrl: './matrix.component.html',
  styleUrl: './matrix.component.css',
})
export class MatrixComponent {
  private matrixService = inject(MatrixService);
  private router = inject(Router);

  size = signal(3);
  matrix = signal<number[][]>(this.generarMatrizVacia(this.size()));
  result = signal<MatrixResponse | null>(null);
  isLoading = signal(false);
  errorMessage = signal('');

  generarMatrizVacia(n: number): number[][] {
    return Array(n)
      .fill(0)
      .map(() => Array(n).fill(0));
  }

  onSizeChange(newSize: number) {
    this.size.set(newSize);
    this.matrix.set(this.generarMatrizVacia(newSize));
    this.result.set(null);
  }

  //Estoy optimizando los valores guardados en memoria con signal update y recorriendo toda la matriz de vuelta
  updateValue(row: number, col: number, val: string) {
    const value = Number(val);

    this.matrix.update((currentMatrix) =>
      currentMatrix.map((r, rowIndex) =>
        rowIndex === row ? r.map((c, colIndex) => (colIndex === col ? value : c)) : r,
      ),
    );
  }

  sendMatrix() {
    this.isLoading.set(true);

    this.matrixService.processMatrix(this.matrix()).subscribe({
      next: (data) => {
        this.isLoading.set(false);

        this.result.set(data);
      },
      error: (error) => {
        this.isLoading.set(false);
        this.errorMessage.set('Error al procesar la matriz');
      },
    });
  }
}
