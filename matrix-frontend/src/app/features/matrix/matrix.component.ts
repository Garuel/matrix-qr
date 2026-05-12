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

  updateValue(row: number, col: number, val: string) {
    const value = Number(val);
    const current = this.matrix();
    current[row][col] = value;
    this.matrix.set([...current]);
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
        alert('Error al procesar la matriz');
      },
    });
  }
}
