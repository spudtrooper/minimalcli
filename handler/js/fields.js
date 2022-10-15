class FieldManager {
  constructor(prefix) {
    this.key_ = prefix + '_FieldManager';
  }

  allFields() {
    let session = this.readFromStorage_();
    return session;
  }

  read(key) {
    console.log('read', key);
    let session = this.readFromStorage_();
    return session[key] || null;
  }

  store(key, val) {
    console.log('store', key, val);
    let session = this.readFromStorage_();
    session[key] = val;
    localStorage[this.key_] = JSON.stringify(session);
  }

  delete(key) {
    console.log('delete', key);
    let session = this.readFromStorage_();
    delete session[key];
    localStorage[this.key_] = JSON.stringify(session);
  }

  clear() {
    console.log('clear');
    localStorage[this.key_] = '{}';
  }

  readFromStorage_() {
    return JSON.parse(localStorage[this.key_] || '{}');
  }
}