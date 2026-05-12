import { Injectable, signal, computed } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { tap } from 'rxjs';
import { environment } from '../../../../environments/environment';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private readonly API_AUTH_URL = `${environment.apiUrl}/auth`;

  token = signal<string | null>(localStorage.getItem('token'));

  estaLoggeado = computed(() => !!this.token());

  constructor(private http: HttpClient) {}

  login(usuario: string, password: string) {
    return this.http
      .post<{ token: string }>(`${this.API_AUTH_URL}/login`, { usuario, password })
      .pipe(
        tap((res) => {
          localStorage.setItem('token', res.token);
          this.token.set(res.token);
        }),
      );
  }

  logout() {
    localStorage.removeItem('token');
    this.token.set(null);
  }
}
