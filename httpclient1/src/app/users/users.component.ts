import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-users',
  template: `
    
      <h1 *ngFor="let u of users" (click)="go(u.id)"> {{u.name}} </h1>

  `,
  styles: [
  ]
})
export class UsersComponent implements OnInit {
  users=[
    {id:"112223",name:"Poppy Singh", },
    {id:"114545",name:"Pogli Singh", },
    {id:"145675",name:"Bebo Singh", },
    {id:"357675",name:"Pogli Aurat", },
  ]
  constructor(private router:Router) { 

  }

  ngOnInit(): void {
  
  }
  go(id:string){
      this.router.navigate(['/user',id])
  }
}
