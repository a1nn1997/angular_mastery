import { Component, EventEmitter, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http'

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  authenticated = false;
  authEmitter = new EventEmitter<boolean>();

  constructor(private http: HttpClient) {
  }

  ngOnInit(): void {
    
    
    this.authEmitter.subscribe(
      (auth: boolean) => {
        this.authenticated = auth;
      }
    );
  }

  logout(): void {
    this.http.post('http://localhost:8069/api/logout', {}, {withCredentials: true})
      .subscribe(() => this.authenticated = false);
  }

}
