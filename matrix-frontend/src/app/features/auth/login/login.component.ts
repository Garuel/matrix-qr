import { Component, inject, model, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../../core/services/auth/auth.service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './login.component.html',
})
export class LoginComponent {
  private authService = inject(AuthService);
  private router = inject(Router);

  usuario = model('');
  password = model('');
  errorMessage = signal('');
  isLoading = signal(false);

  onSubmit() {
    if (this.isLoading()) return;
    if (!this.usuario() || !this.password()) {
      this.errorMessage.set('Por favor, completa todos los campos');
      return;
    }

    this.isLoading.set(true);
    this.errorMessage.set('');

    this.authService.login(this.usuario(), this.password()).subscribe({
      next: () => {
        this.router.navigate(['/matrix']);
      },
      error: (err) => {
        this.isLoading.set(false);
        this.errorMessage.set('Credenciales inválidas. Intenta de nuevo.');
      },
    });
  }
}
