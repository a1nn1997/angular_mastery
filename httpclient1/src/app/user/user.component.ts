import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, ParamMap} from '@angular/router'

@Component({
  selector: 'app-user',
  template: `
<    <h2> you requested for user with id {{userid}} </h2>
>  `,
  styles: [
  ]
})
export class UserComponent implements OnInit {

  userid!:string | null;

  constructor(private route:ActivatedRoute) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params:ParamMap)=>{
      this.userid=params.get('id')
    })
  }

}
