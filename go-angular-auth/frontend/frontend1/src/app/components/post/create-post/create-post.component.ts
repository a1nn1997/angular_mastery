import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { throwError } from 'rxjs';
import { PostService } from '../../shared/post.service';
import { SubredditModel } from '../../subreddit/subreddit-response';
import { SubredditService } from '../../subreddit/subreddit.service';
import { CreatePostPayload } from './create-post.payload';

@Component({
  selector: 'app-create-post',
  templateUrl: './create-post.component.html',
  styleUrls: ['./create-post.component.css']
})
export class CreatePostComponent implements OnInit {
 createPostForm: FormGroup;
  postPayload: CreatePostPayload;
  subreddits: Array<SubredditModel>;

  constructor(private router: Router, private postService: PostService,
    private subredditService: SubredditService) {
    this.postPayload = {
      title: '',
      url: '',
      description: '',
      subredditName: ''
    }
  }

  ngOnInit() {
    this.createPostForm = new FormGroup({
      title: new FormControl('', Validators.required),
      subredditName: new FormControl('', Validators.required),
      url: new FormControl('', Validators.required),
      description: new FormControl('', Validators.required),
    });
    this.subredditService.getAllSubreddits().subscribe((data) => {
      this.subreddits = data;
    }, error => {
      throwError(error);
    });
  }

  createPost() {
    console.log()
    this.postPayload.title = this.createPostForm.get('title').value;
    this.postPayload.subredditName = this.createPostForm.get('subredditName').value;
    this.postPayload.url = this.createPostForm.get('url').value;
    this.postPayload.description = this.createPostForm.get('description').value;

    this.postService.createPost(this.postPayload).subscribe((data) => {
      this.router.navigateByUrl('/');
    }, error => {
      throwError(error);
    })
  }

  discardPost() {
    this.router.navigateByUrl('/');
  }
}
