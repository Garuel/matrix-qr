import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../../environments/environment';
import { MatrixResponse } from '../../domain/dto/matrix/matrix.dto';

@Injectable({ providedIn: 'root' })
export class MatrixService {
  private http = inject(HttpClient);
  private readonly API_URL = `${environment.apiUrl}/matrix`;

  processMatrix(data: number[][]) {
    return this.http.post<MatrixResponse>(`${this.API_URL}/process`, { matrix: data });
  }
}
