import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  constructor(private router:Router) { }

  ngOnInit(): void {
  }
  register(regForm:NgForm) {
    console.log(regForm.value);
    setTimeout(()=>{regForm.reset()},500);
    this.router.navigate(['/login']);
  }
  

}
