import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { throwError } from 'rxjs';
import { CommentPayload } from '../../comment/comment-payload';
import { CommentService } from '../../comment/comment.service';
import { Emitters } from '../../emitters/emitters';
import { PostModel } from '../../shared/post-model';
import { PostService } from '../../shared/post.service';

@Component({
  selector: 'app-view-post',
  templateUrl: './view-post.component.html',
  styleUrls: ['./view-post.component.css']
})
export class ViewPostComponent implements OnInit {

 postId: string;
  post: PostModel;
  commentForm: FormGroup;
  commentPayload: CommentPayload;
  comments: CommentPayload[];
  //username: string="";


  constructor(private postService: PostService,private http: HttpClient, private activateRoute: ActivatedRoute,
    private commentService: CommentService, private router: Router) {
    this.postId = this.activateRoute.snapshot.params['id'];

    this.commentForm = new FormGroup({
      text: new FormControl('', Validators.required)
    });
    this.commentPayload = {
      text: '',
      postId: this.postId
    };
  }

  ngOnInit(): void {
   /* this.http.get('http://localhost:8069/api/user', {withCredentials: true}).subscribe( (res: any) => {
      var fn=String(res.first_name)
      var ln=String(res.last_name)  
      this.username=fn+" "+ln;
       Emitters.authEmitter.emit(true);
      },
      err => {
        this.username = 'temp user';
        Emitters.authEmitter.emit(false);
      }
    );*/
    this.getPostById();
    this.getCommentsForPost();
   // console.log(this.username)
  }

  postComment() {
    this.commentPayload.text = this.commentForm.get('text').value;
    this.commentService.postComment(this.commentPayload).subscribe(data => {
      this.commentForm.get('text').setValue('');
      this.getCommentsForPost();
    }, error => {
      throwError(error);
    })
  }

  private getPostById() {
    this.postService.getPost(this.postId).subscribe(data => {
      this.post = data;
    }, error => {
      throwError(error);
    });
  }

  private getCommentsForPost() {
    this.commentService.getAllCommentsForPost(this.postId).subscribe(data => {
      this.comments = data;
    }, error => {
      throwError(error);
    });
  }
}
