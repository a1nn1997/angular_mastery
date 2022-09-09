import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-edituser',
  templateUrl: './edituser.component.html',
  styleUrls: ['./edituser.component.css']
})
export class EdituserComponent implements OnInit {
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

  password= new FormControl("",[
    Validators.minLength(6),
    
  ]);
  
  ngOnInit(): void {
  
    this.form = this.formBuilder.group({
      username: '',
      first_name: '',
      last_name: '',
      phone:'',
      email: '',
      password: this.password,
    });
  }

  submit(): void {
    //console.log(regForm.value)
    console.log(this.form.getRawValue().username)
    var loginform={}
    loginform={
      "username":this.form.getRawValue().username,
      "email":this.form.getRawValue().email,
      "password":this.form.getRawValue().password,
      "first_name":this.form.getRawValue().first_name,
      "last_name":this.form.getRawValue().last_name,
      "phone":this.form.getRawValue().phone,
    }
    this.http.post('http://localhost:8069/api/edituser', loginform, {withCredentials: true})
      .subscribe(() => this.router.navigate(['/']));
  }

}
