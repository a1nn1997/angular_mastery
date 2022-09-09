import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  form: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private http: HttpClient,
    private router: Router,
    
  ) {
  }

  email= new FormControl("",[
    Validators.email,
    
  ]);

  password1= new FormControl("",[
    Validators.minLength(6),
    
  ]);
  
  password2= new FormControl("",[
    Validators.minLength(6),
  ]);


  pass_match:boolean=this.password1==this.password2;

  ngOnInit(): void {
  
    this.form = this.formBuilder.group({
      username: '',
      first_name: '',
      last_name: '',
      phone:'',
      email: '',
      password1: this.password1,
      password2: this.password2,
      pass_match: this.pass_match
    });
  }

  submit(): void {
    //console.log(regForm.value)
    console.log(this.form.getRawValue().username)
    var loginform={}
    loginform={
      "username":this.form.getRawValue().username,
      "email":this.form.getRawValue().email,
      "password":this.form.getRawValue().password1,
      "first_name":this.form.getRawValue().first_name,
      "last_name":this.form.getRawValue().last_name,
      "phone":this.form.getRawValue().phone,
    }
    this.http.post('http://localhost:8069/api/register', loginform)
      .subscribe(() => this.router.navigate(['/login']));
  }
}
