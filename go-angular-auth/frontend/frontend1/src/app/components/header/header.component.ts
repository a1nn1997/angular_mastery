import { Component,  OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { Emitters } from '../emitters/emitters';
import { Router } from '@angular/router';
import { faUser } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
authenticated = false;
  username: string="";
    faUser = faUser;

  constructor(private http: HttpClient,private router: Router) {
  }

  ngOnInit(): void {
    Emitters.authEmitter.subscribe(
      (auth: boolean) => {
        this.authenticated = auth;
      },
      //username to be taken
    this.http.get('http://localhost:8069/api/user', {withCredentials: true}).subscribe( (res: any) => {
      var fn=String(res.first_name)
      var ln=String(res.last_name)  
      this.username=fn+" "+ln;
       Emitters.authEmitter.emit(true);
      },
      err => {
        this.username = 'please fill your first name and last name';
        Emitters.authEmitter.emit(false);
      }
    )
    );
    } 
 goToUserProfile() {
    this.router.navigateByUrl('/user-profile/' + this.username);
  }
  logout(): void {
    this.http.post('http://localhost:8069/api/logout', {}, {withCredentials: true})
      .subscribe(() => this.authenticated = false);
  }
}
