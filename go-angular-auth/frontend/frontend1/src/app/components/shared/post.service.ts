import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CreatePostPayload } from '../post/create-post/create-post.payload';
import { PostModel } from './post-model';

@Injectable({
  providedIn: 'root'
})
export class PostService {
constructor(
    private http: HttpClient
    ){}
  
    getAllPosts(): Observable<Array<PostModel>> {
      return this.http.get<Array<PostModel>>('http://localhost:8069/api/post', {withCredentials: true});
    }
  
    createPost(postPayload: CreatePostPayload): Observable<any> {
      return this.http.post('http://localhost:8069/api/post', postPayload, {withCredentials: true});
    }
  
    getPost(id: string): Observable<PostModel> {
      return this.http.get<PostModel>('http://localhost:8069/api/post/' + id, {withCredentials: true});
    }
  
    getAllPostsByUser(name: string): Observable<PostModel[]> {
      return this.http.get<PostModel[]>('http://localhost:8069/api/postbyuser/' , {withCredentials: true});
    }
}
