import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { Emitters } from './components/emitters/emitters';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'frontend';
   message='';
  constructor(
    private http:HttpClient
  ){}
  ngOnInit() {
    this.http.get('http://localhost:8069/api/user', {withCredentials: true}).subscribe(
      (res: any) => {
        var fn=String(res.first_name)
        var ln=String(res.last_name)
        this.message = `Hello ${fn[0].toUpperCase()+fn.substring(1)} ${ln[0].toUpperCase()+ln.substring(1)}
        `;
        Emitters.authEmitter.emit(true);
      },
      err => {
        this.message = '';
        Emitters.authEmitter.emit(false);
      }
    );
  }

}
