import { HttpClient } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { faComments } from '@fortawesome/free-solid-svg-icons';
import { Emitters } from '../../emitters/emitters';
import { PostModel } from '../post-model';

@Component({
  selector: 'app-post-title',
  templateUrl: './post-title.component.html',
  styleUrls: ['./post-title.component.css']
})
export class PostTitleComponent implements OnInit {


  faComments = faComments;
  @Input() posts: PostModel[];

  constructor(private router: Router,private http: HttpClient) { }

  ngOnInit(): void {
   
  }

  goToPost(id: number): void {
    this.router.navigateByUrl('/view-post/' + id);
  }
}
